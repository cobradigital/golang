package router

import (
	"fmt"
	"net/http"
	"time"

	"../flags"
	"../response"

	"github.com/gorilla/mux"

	"../configs"
	"../handler"
	"../loggers"
	"../middleware"
)

// Logger
var log = loggers.Get()

// A Router extends gorilla mux.Router functionality to handle RESTFunc
type Router struct {
	*mux.Router
}

// Init auth middleware
var Auth = middleware.Auth()

// Init start time
var startTime time.Time

// HandleREST authenticated request and execute RESTFunc type handler
func (r Router) HandleREST(path string, fn handler.RESTFunc, purpose string) *mux.Route {
	var h http.Handler
	if purpose != flags.ACLEveryone {
		h = Auth(fn, purpose)
	} else {
		h = fn
	}
	return r.NewRoute().Path(path).Handler(h)
}

// New creates new router instance and configure api routing by calling routeAPI() function
func New(start time.Time) Router {
	// Set start time
	startTime = start
	// Create new router
	r := Router{mux.NewRouter()}
	// Get base url
	baseUrl := configs.MustGetString("server.base_url")
	log.Infof("API Base URL: %s", baseUrl)
	// Init api router
	a := Router{r.PathPrefix(baseUrl).Subrouter()}
	routeAPI(a)
	// Set error handler
	r.NotFoundHandler = handler.RESTFunc(handler.NotFound)
	r.MethodNotAllowedHandler = handler.RESTFunc(handler.MethodNotAllowed)
	// Set main handler
	r.HandleREST(baseUrl, GetAppStatus, flags.ACLEveryone).Methods("GET")
	// Return main router
	return r
}

func GetAppStatus(_ *http.Request) (*response.Success, error) {
	body := AppStatus{
		BuildVersion: flags.AppVersion,
		Uptime:       fmt.Sprintf("%s", time.Since(startTime)),
	}
	return response.NewSuccess(&body), nil
}

type AppStatus struct {
	BuildVersion string `json:"build_version"`
	Uptime       string `json:"uptime"`
}
