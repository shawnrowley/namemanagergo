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
	// PersonController struct
	PersonController struct {
		session *mgo.Session
	}
)

// NewPersonController creates
func NewPersonController(s *mgo.Session) *PersonController {
	return &PersonController{s}
}

func getURLJson(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(target)
}

// GetGenderReport get Person Report
func (pc PersonController) GetGenderReport(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var reports []models.Report
	var names []string

	err := pc.session.DB("namemanager").C("Person").Find(nil).Distinct("firstName", &names)
	if err != nil {
		w.WriteHeader(404)
		return
	}

	for _, name := range names {
		report := new(models.Report)
		gender := new(models.Gender)
		getURLJson("https://gender-api.com/get?key=drdTuwqCXVAMYkReEo&name="+name, gender)

		count, err := pc.session.DB("namemanager").C("Person").Find(bson.M{"firstName": name}).Count()
		if err != nil {
			w.WriteHeader(404)
			return
		}
		report.Instances = count
		report.Gender = gender.Gender
		report.Accuracy = gender.Accuracy
		report.Name = name
		reports = append(reports, *report)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := json.NewEncoder(w).Encode(reports); err != nil {
		panic(err)
	}
}

// GetPerson gets Person by Id
func (pc PersonController) GetPerson(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
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

// GetAllPersons gets all persons
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
	person := models.Person{}
	json.NewDecoder(r.Body).Decode(&person)
	person.Id = bson.NewObjectId()
	pc.session.DB("namemanager").C("Person").Insert(person)
	personjson, _ := json.Marshal(person)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", personjson)
}

// UpdatePerson updates Person resource
func (pc PersonController) UpdatePerson(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	id := p.ByName("id")
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}
	oid := bson.ObjectIdHex(id)

	person := models.Person{}
	json.NewDecoder(r.Body).Decode(&person)

	selector := bson.M{"_id": oid}
	updates := bson.M{"$set": bson.M{"firstName": person.FirstName, "lastName": person.LastName, "ipAddress": person.IpAddress, "country": person.Country}}
	if err := pc.session.DB("namemanager").C("Person").Update(selector, updates); err != nil {
		w.WriteHeader(404)
		return
	}

	if err := pc.session.DB("namemanager").C("Person").FindId(oid).One(&person); err != nil {
		w.WriteHeader(404)
		return
	}
	personjson, _ := json.Marshal(person)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", personjson)
}

// DeletePerson removes an existing Person resource
func (pc PersonController) DeletePerson(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}
	oid := bson.ObjectIdHex(id)

	// Delete person
	if err := pc.session.DB("namemanager").C("Person").RemoveId(oid); err != nil {
		w.WriteHeader(404)
		return
	}
	w.WriteHeader(200)
}
