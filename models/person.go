package models

import "gopkg.in/mgo.v2/bson"

type (
    // Person struct
    Person struct {
        Id     		  bson.ObjectId 	`json:"id" bson:"_id"`
        FirstName   string 		`json:"firstName" bson:"firstName"`
        LastName	  string 		`json:"lastName" bson:"lastName"`
        IpAddress   string 		`json:"ipAddress" bson:"ipAddress"`
      	Country 	  string 		`json:"country" bson:"country"`
    }
)

type Persons []Person
