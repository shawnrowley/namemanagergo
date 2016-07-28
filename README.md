# namemanagergo

====================
Author: Shawn A. Rowley  
Source: <https://github.com/shawnrowley/namemanagergo/>  

Overview
-----------

Create an app that allows you to manage names. Basic CRUD operation. Also include a report which will display all the names, how many times each name has been entered, and query this public API to also show the likely gender and confidence level.

Public API: https://gender-api.com/en/api-docs

System Requirements
-----------

Golang 
MongoDB

Developer Notes
-----------

Digging Golang, like a good book couldn't put it down. Good resources online. Basic CRD :) GTG, forgot update

Completed the baseline implementation to include the report. 

Added update functionality now CRUD GTG.


RESTful API
-----------

    r.GET("/person", pc.GetAllPersons)
    r.GET("/person/:id", pc.GetPerson)
    r.POST("/person", pc.CreatePerson)
    r.DELETE("/person/:id", pc.DeletePerson)
    r.PUT("/person/:id", pc.UpdatePerson)
    r.GET("/report", pc.GetGenderReport);

Usage
--------------
Create: 
	curl -XPOST -H 'Content-Type: application/json' -d '{"firstname": "Shawn", "lastname": "Rowley", "ipAddress":
	"127.0.0.1", "country": "US"}' http://localhost:3000/person

Get All Person
	curl http://localhost:3000/person/

Get Person by id
	curl http://localhost:3000/person/579529c1692044391a3aa8fd  (GET)
	
Update:	
	curl -XPUT -H 'Content-Type: application/json' -d '{"firstname": "Shawn", "lastname": "Rowley", "ipAddress":
	"127.0.0.1", "country": "US"}' http://localhost:3000/person/579529c1692044391a3aa8fd  
	
Delete:
        curl http://localhost:3000/person/579529c1692044391a3aa8fd  (DELETE)
        
Report:
        curl http://localhost:3000/report       

    
Technologies
-----------

Golang
MongoDB


Project/Server
-----------

Atom
Server: Golang (julienschmidt/httprouter)


Deployment
-----------

	
Enhancements/Ideas
-----------	

	Try gorilla 
	Needs a UI

Code References 
------------


