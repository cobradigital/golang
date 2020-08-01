package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	"./configs"
	"./crypto"
	"./flags"
	"./loggers"
	"./repositories"
	"./router"
	"./services"
	"./util"
	"github.com/rs/cors"
)

var log = loggers.Get()

func main() {

	showVersion := flag.Bool("version", false, "Version")

	// Get starting time
	start := time.Now()

	flag.Parse()

	if *showVersion {
		printVersion()
		os.Exit(0)
	}

	configFile := flag.String("c", "config.yml", "config.yml")
	if *configFile == "" {
		fmt.Println("Configuration file is not specified. Use -c config.yml to specify one")
		os.Exit(1)
	}
	configs.Init(*configFile)

	r := router.New(start)
	// Start server

	repositories.Init()

	util.Init()

	services.Init()

	// Initiate crypto
	crypto.Init()

	// Where ORIGIN_ALLOWED is like `scheme://dns[:port]`, or `*` (insecure)
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},                                     // All origins
		AllowedMethods: []string{"GET", "HEAD", "POST", "PUT", "OPTIONS"}, // Allowing only get, just an example
	})

	log.Infof("Boot time: %s", time.Since(start))
	log.Fatal(http.ListenAndServe(getPort(), c.Handler(r)))
}

// getPort retrieve port from config
func getPort() string {
	port := configs.MustGetString("server.port")
	return ":" + port
}

func printVersion() {
	fmt.Printf("%s version %s, build %s\n", flags.AppName, flags.AppVersion, flags.AppCommitHash)
}
