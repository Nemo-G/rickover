package main

import (
	"log"

	"rickover/models/db"
	"rickover/setup"
	"rickover/test"
)

func main() {
	if err := setup.DB(db.DefaultConnection, 1); err != nil {
		log.Fatal(err)
	}
	if err := test.TruncateTables(nil); err != nil {
		log.Fatal(err)
	}
}
