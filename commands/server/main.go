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

func configure() (http.Handler, error) {
	dbConns, err := config.GetInt("PG_SERVER_POOL_SIZE")
	if err != nil {
		log.Printf("PG_SERVER_POOL_SIZE error: %s. Defaulting to 10", err)
		dbConns = 10
	}

	if err = setup.DB(db.DefaultConnection, dbConns); err != nil {
		return nil, err
	}

	metrics.Namespace = "rickover.server"
	metrics.Start("web")

	go setup.MeasureActiveQueries(5 * time.Second)

	// If you run this in production, change this user.
	server.AddUser("test", "hymanrickover")
	return server.Get(server.DefaultAuthorizer), nil
}

func main() {
	s, err := configure()
	if err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}
	log.Printf("Listening on port %s\n", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), handlers.LoggingHandler(os.Stdout, s))
}
