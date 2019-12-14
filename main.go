package main

import (
	"http"

	"github.com/gorilla/mux"
	"github.com/reaganiwadha/bookrest/config"
)

var r *mux.Route

func main() {
	config.ThisExportedFunction()
	r := mux.NewRouter()
	http.Handle("/", r)
}
