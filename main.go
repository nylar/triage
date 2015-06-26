package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/nylar/triage/api"
	"github.com/nylar/triage/app"
)

// Version
const (
	MAJOR = 0
	MINOR = 1
	PATCH = 0
)

var version = fmt.Sprintf("%d.%d.%d", MAJOR, MINOR, PATCH)

func main() {
	fmt.Printf("Triage, v%s\n", version)

	// Create connection to the database
	db, err := app.Connect("postgres", "", "triage")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer db.Close()

	// Initialse a new router and turn on missing trailing slash redirect
	m := mux.NewRouter()
	m.StrictSlash(true)

	// Setup API routes
	api.Routes(m, db)

	// Setup a route for the public facing site
	m.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))

	http.Handle("/", m)

	log.Println("Serving on port :3030.")
	http.ListenAndServe(":3030", nil)
}
