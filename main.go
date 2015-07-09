package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
	"github.com/nylar/triage/api"
	"github.com/nylar/triage/app"
)

// Version
const (
	Major = 0
	Minor = 1
	Patch = 0
)

var version = fmt.Sprintf("%d.%d.%d", Major, Minor, Patch)

func main() {
	fmt.Printf("Triage, v%s\n", version)

	// Create connection to the database
	db, err := app.Connect("postgres", "", "triage")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer db.Close()

	// Initialse a new router
	router := httprouter.New()

	// Setup API routes
	api.Routes(router, db)

	log.Println("Serving on port :3030.")
	http.ListenAndServe(":3030", router)
}
