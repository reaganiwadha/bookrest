package dao

import (
	"log"

	"gopkg.in/mgo.v2"
)

var db *mgo.Database

//BookDAO ... exported
type BookDAO struct {
	Server   string
	Database string
}

const ()

//Connect ... Connects to the db
func (b *BookDAO) Connect() {
	session, err := mgo.Dial(b.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(b.Database)
}

func main() {

}
