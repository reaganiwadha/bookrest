//github.com/reaganiwadha

// THIS IS A REALLY REALLY TERRIBLE CODE I AM REALLY SORRY

package main

//I dont understand submodules in go... yet
//So here's to a one filer

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/BurntSushi/toml"
	"github.com/common-nighthawk/go-figure"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Book ... the type for books on the db
type Book struct {
	ISBN          string
	Title         string
	Author        string
	Status        string
	Publisher     string
	Year          int //im sorrya
	IssueCount    int
	ViewCount     int
	CurrentIssuer string
}

type NewBook struct {
	Title     string
	Author    string
	Isbn      string
	Status    string
	Publisher string
	Year      int
}

// BookrestConfig ... the config for this server
type BookrestConfig struct {
	Port                     int
	Database                 string
	DatabaseConnectionString string
	BooksCollection          string
}

// Issuer ... who's issuing the book?
type Issuer struct {
	Issuer string
}

var config BookrestConfig
var configPath string = "config.toml"

var ctx context.Context

var session mgo.Session
var db *mgo.Database

//AllBooksEndpoint ... returns all the books
func AllBooksEndpoint(w http.ResponseWriter, r *http.Request) {
	var books []Book
	err := db.C(config.BooksCollection).Find(bson.M{}).All(&books)
	if err != nil {
		log.Fatal(err)
	}
	
	RespondWithJSON(w, http.StatusOK, books)
}

//IssueEndpoint ... To issue a book
func IssueEndpoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var book Book
	params := mux.Vars(r)
	data := db.C(config.BooksCollection).Find(bson.M{"isbn": params["isbn"]})
	err := data.One(&book)
	if err != nil {
		log.Fatal(err)
		RespondWithError(w, http.StatusBadRequest, "Error on db")
		return
	}

	var issuer Issuer
	if err := json.NewDecoder(r.Body).Decode(&issuer); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if book.Status != "issued" {
		change := mgo.Change{
			Update:    bson.M{"$inc": bson.M{"issuecount": 1}, "$set": bson.M{"status": "issued", "currentissuer": issuer.Issuer}},
			ReturnNew: true,
		}
		_, err := data.Apply(change, &book)
		if err != nil {
			log.Fatal(err)
		}
		RespondWithJSON(w, http.StatusOK, book)
	} else {
		RespondWithError(w, http.StatusBadRequest, "Book already issued")
	}
}

//DeleteIssuerEndpoint ... delete the current issuer on the book
func DeleteIssuerEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var book Book
	data := db.C(config.BooksCollection).Find(bson.M{"isbn": params["isbn"]})
	err := data.One(&book)
	if err != nil {
		log.Fatal(err)
		RespondWithError(w, http.StatusBadRequest, "Error on db")
		return
	}

	change := mgo.Change{
		Update:    bson.M{"$set": bson.M{"status": "available", "currentissuer": ""}},
		ReturnNew: true,
	}
	_, errr := data.Apply(change, &book)
	if errr != nil {
		log.Fatal(err)
		RespondWithError(w, http.StatusBadRequest, "Error on db")
		return
	}
	RespondWithOK(w)
}

//FindBookEndpoint ... Find book based on ISBN
func FindBookEndpoint(w http.ResponseWriter, r *http.Request) {
	var book []Book
	params := mux.Vars(r)

	change := mgo.Change{
		Update:    bson.M{"$inc": bson.M{"viewcount": 1}},
		ReturnNew: true,
	}

	data := db.C(config.BooksCollection).Find(bson.M{"isbn": params["isbn"]})

	data.Apply(change, &book)
	data.All(&book)
	RespondWithJSON(w, http.StatusOK, book)
}

//TopBookEndpoint ... returns the book that is most viewed
func TopBookEndpoint(w http.ResponseWriter, r *http.Request) {
	var books []Book
	err := db.C(config.BooksCollection).Find(bson.M{}).Sort("-viewcount").Limit(1).All(&books)
	if err != nil {
		log.Fatal(err)
	}
	
	RespondWithJSON(w, http.StatusOK, books)
}

//MostIssuedBookEndpoint ... returns the book that is most issued
func MostIssuedBookEndpoint(w http.ResponseWriter, r *http.Request) {
	var books []Book
	err := db.C(config.BooksCollection).Find(bson.M{}).Sort("-issuecount").Limit(1).All(&books)
	if err != nil {
		log.Fatal(err)
	}
	RespondWithJSON(w, http.StatusOK, books)
}

//CreateBookEndpoint ... Creates a new book on the database
func CreateBookEndpoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var book NewBook
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
	}
	db.C(config.BooksCollection).Insert(book)
	RespondWithOK(w)
}

//RespondWithJSON ... A method to respond with json data
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

//RespondWithOK ... A method to return a 200 code
func RespondWithOK(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
}

//RespondWithError ... Throws an error
func RespondWithError(w http.ResponseWriter, code int, msg string) {
	RespondWithJSON(w, code, map[string]string{"error": msg})
}

func init() {
	//prints a cool ascii art because why not (im a terrible person for prioritizing ascii art
	//and not good code)

	figure.NewFigure("bookrest", "computer", false).Print()
	fmt.Print("\n\n")

	//Loads the TOML file for config
	if _, err := toml.DecodeFile(configPath, &config); err != nil {
		log.Fatal(err)
	}

	//Dials the database
	session, err := mgo.Dial(config.DatabaseConnectionString)
	if err != nil {
		log.Fatal(err)
	}

	db = session.DB(config.Database)
}

func main() {
	r := mux.NewRouter()

	fmt.Println("Running on port", config.Port)

	r.HandleFunc("/books", AllBooksEndpoint).Methods("GET")
	r.HandleFunc("/books", CreateBookEndpoint).Methods("POST")
	r.HandleFunc("/books/top", TopBookEndpoint).Methods("GET")
	r.HandleFunc("/books/mostissued", MostIssuedBookEndpoint).Methods("GET")

	r.HandleFunc("/books/{isbn}", FindBookEndpoint).Methods("GET")
	r.HandleFunc("/issue/{isbn}", IssueEndpoint).Methods("PUT")
	r.HandleFunc("/issue/{isbn}", DeleteIssuerEndpoint).Methods("DELETE")

	http.Handle("/", r)
	http.ListenAndServe(":"+strconv.Itoa(config.Port), nil)
}
