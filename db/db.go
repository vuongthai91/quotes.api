package db

import (
	"log"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Author struct {
	ID          bson.ObjectId `json:"_id" bson:"_id"`
	Name        string        `json:"name" bson:"name"`
	Nationality string        `json:"nationality" bson:"nationality"`
	Profession  string        `json:"profession" bson:"profession"`
	Birthday    string        `json:"birthday" bson:"birthday"`
	Diedday     string        `json:"diedday" bson:"diedday"`
	TotalView   int64         `json:"total_view" bson:"total_view"`
	GroupLetter string        `json:"group_letter" bson:"group_letter"`
	CreatedDate time.Time     `json:"created_date" bson:"created_date"`
	UpdatedDate time.Time     `json:"updated_date" bson:"updated_date"`
}

var db *mgo.Database

func init() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	db = session.DB("cleverquote")
}

func collection() *mgo.Collection {
	return db.C("Author")
}

// GetAll returns all items from the database.
func GetAll() ([]Author, error) {
	res := []Author{}

	if err := collection().Find(nil).All(&res); err != nil {
		return nil, err
	}

	return res, nil
}

// GetOne returns a single item from the database.
func GetLetter(letter string) (*Author, error) {
	res := Author{}

	if err := collection().Find(bson.M{"group_letter": letter}).One(&res); err != nil {
		return nil, err
	}

	log.Println(letter)
	return &res, nil
}

func GetOne(name string) (*Author, error) {
	res := Author{}

	if err := collection().Find(bson.M{"name": name}).One(&res); err != nil {
		return nil, err
	}

	return &res, nil
}

// Save inserts an item to the database.
func Save(item Author) error {
	return collection().Insert(item)
}

// Remove deletes an item from the database
func Remove(id string) error {
	return collection().Remove(bson.M{"_id": id})
}
