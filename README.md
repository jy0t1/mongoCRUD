# mongoCRUD
CRUD operation on MongoDB using Go. This project is developed in Windows10.

Folder structure
$GOPATH\src\github.com\jy0t1\mongoCRUD
Go files under mongoCRUD/
        main.go                # main file containing routes/API calls
        db.go                  # for Database connection
        controllers.go         # Handelers for routes

To build:
mongoCRUD>go build main.go db.go controllers.go

To run => mongoCRUD>main
This opens the service.

In browser => http://localhost:8080/getall
or use postman to test CRUD functionalities. MongoDB is in AWS.
