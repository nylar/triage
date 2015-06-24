package api

import "github.com/gorilla/mux"

var router *mux.Router

func init() {
	router = mux.NewRouter()
	router.StrictSlash(true)
}
