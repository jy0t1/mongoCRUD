package main

import (
"github.com/gorilla/mux"
"log"
"fmt"
"net/http"
//"go.mongodb.org/mongo-driver/mongo"
//"go.mongodb.org/mongo-driver/mongo/options"
//"context"
// "time"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome home! Jyoti is building a RESTful JSON API with GOlang for CRUD operation on MongoDB collection.")
	fmt.Fprintln(w, "Pass URL parameters to route to proper CRUD operation.")
}

func main() {
	fmt.Println("mongoCRUD => main.go 101")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	
	router.HandleFunc("/create", createBook).Methods("POST")
	router.HandleFunc("/getall", getAllBooks).Methods("GET")
	router.HandleFunc("/getone", getBook).Methods("GET")
	router.HandleFunc("/update", updateBook).Methods("PATCH")
	router.HandleFunc("/delete", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}


