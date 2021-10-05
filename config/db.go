package config

import (
	"fmt"

	"gopkg.in/mgo.v2"
)

// Database
var DB *mgo.Database

// Collections
var Books *mgo.Collection

func init() {
	// envErr := godotenv.Load()
	// if envErr != nil {
	// 	panic(envErr)
	// }
	s, err := mgo.Dial("mongodb://localhost/bookstore")
	if err != nil {
		panic(err)
	}

	if err = s.Ping(); err != nil {
		panic(err)
	}

	DB = s.DB("bookstore")
	Books = DB.C("books")

	fmt.Println("Connected to MongoDB..")
}
