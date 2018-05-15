package main

import (
	"log"
	"net/http"

	"github.com/rs/cors"
	"github.com/rs/rest-layer-mongo"
	"github.com/rs/rest-layer/resource"
	"github.com/rs/rest-layer/rest"
	"github.com/rs/rest-layer/schema"
	"gopkg.in/mgo.v2"
)

var (
	user = schema.Schema{
		Fields: schema.Fields{
			"id":      schema.IDField,
			"created": schema.CreatedField,
			"updated": schema.UpdatedField,
			"name": {
				Required:   true,
				Filterable: true,
				Sortable:   true,
				Validator: &schema.String{
					MaxLen: 150,
				},
			},
		},
	}

	author = schema.Schema{
		Fields: schema.Fields{
			"id":           schema.IDField,
			"created_date": schema.CreatedField,
			"updated_date": schema.UpdatedField,
			"name": {
				Required:   true,
				Filterable: true,
				Sortable:   true,
				Validator: &schema.String{
					MaxLen: 150,
				},
			},
		},
	}

	quote = schema.Schema{
		Fields: schema.Fields{
			"id":           schema.IDField,
			"created_date": schema.CreatedField,
			"updated_date": schema.UpdatedField,
			"name": {
				Required:   true,
				Filterable: true,
				Sortable:   true,
				Validator: &schema.String{
					MaxLen: 150,
				},
			},
			"author": {
				Required:   true,
				Filterable: true,
				Validator: &schema.Reference{
					Path: "authors",
				},
			},
		},
	}
)

func main() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		log.Fatalf("Can't connect to MongoDB: %s", err)
	}
	db := "cleverquote"

	index := resource.NewIndex()

	index.Bind("authors", author, mongo.NewHandler(session, db, "Author"), resource.DefaultConf)
	index.Bind("quote", quote, mongo.NewHandler(session, db, "Quote"), resource.DefaultConf)

	api, err := rest.NewHandler(index)
	if err != nil {
		log.Fatalf("Invalid API configuration: %s", err)
	}

	http.Handle("/v1/", cors.New(cors.Options{OptionsPassthrough: true}).Handler(http.StripPrefix("/v1/", api)))

	log.Print("Serving API on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
