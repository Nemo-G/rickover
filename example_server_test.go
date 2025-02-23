// Run the rickover server.
//
// All of the project defaults are used. There is one authenticated user for
// basic auth, the user is "test" and the password is "hymanrickover". You will
// want to copy this binary and add your own authentication scheme.
package rickover

import (
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

var serverDbConns int

func init() {
	var err error
	serverDbConns, err = config.GetInt("PG_SERVER_POOL_SIZE")
	if err != nil {
		log.Printf("PG_SERVER_POOL_SIZE error: %s. Set to default 10", err)
		serverDbConns = 10
	}

	metrics.Namespace = "rickover.server"

	// Change this user to a private value
	server.AddUser("test", "hymanrickover")
}

func Example_server() {
	setup.MustSetupDB(db.DefaultConnection, serverDbConns)

	metrics.Start("web")

	go setup.MeasureActiveQueries(5 * time.Second)

	log.Println("Listening on port 9090")
	log.Fatal(http.ListenAndServe(":9090", handlers.LoggingHandler(os.Stdout, server.DefaultServer)))
}
