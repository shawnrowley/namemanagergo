package controllers

import (  
    "encoding/json"
    "fmt"
    "net/http"

    "github.com/julienschmidt/httprouter"
    "github.com/shawnrowley/namemanagergo/models"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson" 
)

type (  
    // PersonController represents the controller for operating on the Person resource
    PersonController struct {
		session *mgo.Session
	}
)

// NewPersonController provides a reference to a PersonController with provided mongo session
func NewPersonController(s *mgo.Session) *PersonController {
	return &PersonController{s}
}

// GetPerson retrieves an individual Person resource
func (pc PersonController) GetPerson(w http.ResponseWriter, r *http.Request, p httprouter.Params) {  
       // Grab id
	id := p.ByName("id")

	// Verify id is ObjectId, otherwise bail
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}

	// Grab id
	oid := bson.ObjectIdHex(id)

	// Stub user
	person := models.Person{}

	// Fetch user
	if err := pc.session.DB("namemanager").C("Person").FindId(oid).One(&person); err != nil {
		w.WriteHeader(404)
		return
	}

	// Marshal provided interface into JSON strpcture
	uj, _ := json.Marshal(person)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", uj)
}

// CreatePerson creates a new Person resource
func (pc PersonController) CreatePerson(w http.ResponseWriter, r *http.Request, p httprouter.Params) {  
    // Stub an Person to be populated from the body
    person := models.Person{}

    // Populate the Person data
    json.NewDecoder(r.Body).Decode(&person)

    // Add an Id
    person.Id = bson.NewObjectId()

    // Write the user to mongo
    pc.session.DB("namemanager").C("Person").Insert(person)	

    // Marshal provided interface into JSON strpcture
     personjson, _ := json.Marshal(person)

    // Write content-type, statuscode, payload
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(201)
    fmt.Fprintf(w, "%s", personjson)
}

// RemovePerson removes an existing Person resource
func (pc PersonController) RemovePerson(w http.ResponseWriter, r *http.Request, p httprouter.Params) {  
    	// Grab id
	id := p.ByName("id")

	// Verify id is ObjectId, otherwise bail
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}

	// Grab id
	oid := bson.ObjectIdHex(id)

	// Remove user
	if err := pc.session.DB("namemanager").C("Person").RemoveId(oid); err != nil {
		w.WriteHeader(404)
		return
	}

	// Write status
	w.WriteHeader(200)
}

