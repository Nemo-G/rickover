package main

import (
	"log"

	"rickover/models/db"
	"rickover/setup"
	"rickover/test"
)

func main() {
	setup.MustSetupDB(db.DefaultConnection, 1)
	if err := test.TruncateTables(nil); err != nil {
		log.Fatal(err)
	}
}
