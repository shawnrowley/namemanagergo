package main

import (
    "net/http"
    "github.com/julienschmidt/httprouter"
    "github.com/shawnrowley/namemanagergo/controllers"
    "gopkg.in/mgo.v2"
)

func getSession() *mgo.Session {
    // Connect to our local mongo
    s, err := mgo.Dial("mongodb://localhost")

    if err != nil {
        panic(err)
    }
    return s
}

func main() {
    r := httprouter.New()
    pc := controllers.NewPersonController(getSession())

    r.GET("/person", pc.GetAllPersons)
    r.GET("/person/:id", pc.GetPerson)
    r.POST("/person", pc.CreatePerson)
    r.DELETE("/person/:id", pc.DeletePerson)

    // Starts server
    http.ListenAndServe("localhost:3000", r)
}
