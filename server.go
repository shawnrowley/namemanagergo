package main

import (
    // Standard library packages
    "net/http"

    // Third party packages
    "github.com/julienschmidt/httprouter"
    "github.com/shawnrowley/namemanagergo/controllers"
    "gopkg.in/mgo.v2"
)

func getSession() *mgo.Session {
    // Connect to our local mongo
    s, err := mgo.Dial("mongodb://localhost")

    // Check if connection error, is mongo running?
    if err != nil {
        panic(err)
    }
    return s
}

func main() {
   // Instantiate a new router
   r := httprouter.New()
    // Get a PersonController instance
    pc := controllers.NewPersonController(getSession())

    // Get a person resource
    r.GET("/person", pc.GetAllPersons)

    r.GET("/person/:id", pc.GetPerson)

    r.POST("/person", pc.CreatePerson)

    r.DELETE("/person/:id", pc.RemovePerson)
    // Fire up the server
    http.ListenAndServe("localhost:3000", r)

}
