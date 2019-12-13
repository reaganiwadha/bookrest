package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/reaganiwadha/bookrest/dao"
	"github.com/reaganiwadha/bookrest/config"
)

var dao = BookDAO{}
var config = Config{}

func init() {
	config.Read()
	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}

func main() {
	r := mux.NewRouter()
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
