// Run the rickover server.
//
// All of the project defaults are used. There is one authenticated user for
// basic auth, the user is "test" and the password is "hymanrickover". You will
// want to copy this binary and add your own authentication scheme.
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"metrics"
	"rickover/config"
	"rickover/models/db"
	"rickover/server"
	"rickover/setup"

	"github.com/gorilla/handlers"
)

func main() {
	dbConns, err := config.GetInt("PG_SERVER_POOL_SIZE")
	if err != nil {
		log.Printf("No PG_SERVER_POOL_SIZE configured: %s. Defaulting to 10", err)
		dbConns = 10
	}
	setup.MustSetupDB(db.DefaultConnection, dbConns)

	metrics.Namespace = "rickover.server"
	metrics.Start("web")
	go setup.MeasureActiveQueries(5 * time.Second)

	port := os.Getenv("PORT")
	if port == "" {
		port = "9090"
	}

	// test acount only work for non production env
	if prod := os.Getenv("Production"); prod != "" {
		server.AddUser("test", "test")
	}
	httpServer := server.Get(server.DefaultAuthorizer)

	log.Printf("Listening on port %s\n", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), handlers.LoggingHandler(os.Stdout, httpServer))
}
