package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"text/template"

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

func indexHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	rootDir, err := os.Getwd()
	if err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(500), 500)
		return
	}

	templatePath := filepath.Join(filepath.Join(rootDir, "public"), "index.html")

	tmpl := template.New("index")
	tmpl = template.Must(template.ParseFiles(templatePath))

	if err := tmpl.Execute(w, nil); err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(500), 500)
		return
	}
}

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

	// Setup a route for the public facing site
	router.GET("/", indexHandler)
	router.ServeFiles("/public/*filepath", http.Dir("public"))

	log.Println("Serving on port :3030.")
	http.ListenAndServe(":3030", router)
}
