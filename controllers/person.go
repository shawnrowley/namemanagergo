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
		PersonController struct {
		session *mgo.Session
	}
)

func NewPersonController(s *mgo.Session) *PersonController {
	return &PersonController{s}
}

func (pc PersonController) GetPerson(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	 id := p.ByName("id")
 // Verify id is ObjectId, otherwise bail
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}
	oid := bson.ObjectIdHex(id)
	person := models.Person{}
	if err := pc.session.DB("namemanager").C("Person").FindId(oid).One(&person); err != nil {
		w.WriteHeader(404)
		return
	}
	personjson, _ := json.Marshal(person)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", personjson)
}

func (pc PersonController) GetAllPersons(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

		var persons []models.Person
		err := pc.session.DB("namemanager").C("Person").Find(nil).All(&persons)
		if err != nil {
			w.WriteHeader(404)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		if err := json.NewEncoder(w).Encode(persons); err != nil {
				panic(err)
		}
}

// CreatePerson creates a new Person resource
func (pc PersonController) CreatePerson(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		// Stub an Person to be populated from the body
		person := models.Person{}
		json.NewDecoder(r.Body).Decode(&person)
		person.Id = bson.NewObjectId()
		pc.session.DB("namemanager").C("Person").Insert(person)
		personjson, _ := json.Marshal(person)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(201)
		fmt.Fprintf(w, "%s", personjson)
}

// RemovePerson removes an existing Person resource
func (pc PersonController) RemovePerson(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}
	oid := bson.ObjectIdHex(id)

	// Remove person
	if err := pc.session.DB("namemanager").C("Person").RemoveId(oid); err != nil {
		w.WriteHeader(404)
		return
	}
	w.WriteHeader(200)
}
