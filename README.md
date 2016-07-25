# namemanagergo

====================
Author: Shawn A. Rowley  
Source: <https://github.com/shawnrowley/namemanagergo/>  


Overview
-----------

Create an app that allows you to manage names. Basic CRUD operation. Also include a report which will display all the names, how many times each name has been entered, and query this public API to also show the likely gender and confidence level.

Public API: https://gender-api.com/en/api-docs


Developer Notes
-----------

Digging Golang, like a good book couldn't put it down. Basic CRUD GTG.

Completed the baseline implementation to include the report.


RESTful API
-----------

    r.GET("/person", pc.GetAllPersons)
    r.GET("/person/:id", pc.GetPerson)
    r.POST("/person", pc.CreatePerson)
    r.DELETE("/person/:id", pc.DeletePerson)
    r.GET("/report", pc.GetGenderReport);

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



