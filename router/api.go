package router

import (
	"../flags"
	"../handler"
)

// routeAPI configure request routing in API. Handlers must be defined in handler package
func routeAPI(r Router) {
	r.HandleREST("/home", handler.Home, flags.ACLEveryone).Methods("POST")
	r.HandleREST("/auth", handler.PostToken, flags.ACLEveryone).Methods("POST")

	// Wamobi
	r.HandleREST("/balance", handler.GetBalance, flags.ACLAuthenticatedUser).Methods("GET")

	// Bot AI
}
