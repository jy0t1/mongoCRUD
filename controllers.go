package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	//"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)
// insertone() will always create lower case elements based on the names given in structure. 
//If field is defined Author_Email then insertone() will create the element as author_email in document
type BookItem struct {
	Id          int64 	`json:id`
	Name 		string  `json:name`
	Author  	string  `json:author`
	AuthorEmail string  `json:authorEmail`
	Published 	string 	`json:published`
	Pages 		int64  	`json:pages`
	Publisher 	string 	`json:publisher`
    IsAvailable bool 	`json:isAvailable`
    Category 	string 	`json:category`
    BindType 	string 	`json:bindType`
    PhotoPath 	string 	`json:photoPath`
}

var bookCollection = db().Database("library_db").Collection("book_records") 

// add new book

func createBook(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json") // for adding Content-type

	// get the values from URL parameters
	ids, ok := r.URL.Query()["id"]
	if !ok || len(ids[0]) < 1 {
        log.Println("Url Param 'id' is missing!")
        return
	}

	name, ok_name := r.URL.Query()["name"]
	if !ok_name || len(name[0]) < 1 {
        log.Println("Url Param 'name' is missing!")
        return
	}

	author, ok_author := r.URL.Query()["author"]
	if !ok_author || len(author[0]) < 1 {
        log.Println("Url Param 'author' is missing!")
        return
	}

	authorEmail, ok_authoremail := r.URL.Query()["authoremail"]
	if !ok_authoremail {
        log.Println("Url Param 'author-email' has issue!")
        return
	}

	published, ok_published := r.URL.Query()["published"]
	if !ok_published {
        log.Println("Url Param 'published' has issue!")
        return
	}

	pages, ok_pages := r.URL.Query()["pages"]
	if !ok_pages {
        log.Println("Url Param 'pages' has issue!")
        return
	}

	publisher, ok_pub := r.URL.Query()["publisher"]
	if !ok_pub {
        log.Println("Url Param 'publisher' has issue!")
        return
	}

	isAvailable, ok_avail := r.URL.Query()["isavailable"]
	if !ok_avail {
        log.Println("Url Param 'isavailable' has issue!")
        return
	}

	category, ok_cat := r.URL.Query()["category"]
	if !ok_cat || len(category[0]) < 1   {
        log.Println("Url Param 'category' is missing!")
        return
	}

	bindType, ok_bind := r.URL.Query()["bindtype"]
	if !ok_bind || len(bindType[0]) < 1   {
        log.Println("Url Param 'bindtype' is missing!")
        return
	}

	photoPath, ok_photo := r.URL.Query()["photopath"]
	if !ok_photo || len(photoPath[0]) < 1   {
        log.Println("Url Param 'photopath' is missing!")
        return
	}
	// get the values from array to local variables
	//strconv.ParseInt(s, 10, 64) to convert to int64
	//inputId, errConv := strconv.Atoi(ids[0])
	inputId, errConv := strconv.ParseInt(ids[0], 10, 64)
	if errConv != nil {
		fmt.Println("in controllers => updatebook")
		fmt.Println(errConv)
	}
	inputPages, errConvPages := strconv.ParseInt(pages[0], 10, 64)
	if errConvPages != nil {
		fmt.Println(errConv)
	}
	inputName := name[0]
	inputAuthor := author[0]
	inputAuthorEmail := authorEmail[0]
	inputPublished := published[0]
	inputPublisher := publisher[0]
	
	inputIsAvail, errBool := strconv.ParseBool(isAvailable[0])
	if errBool != nil {
	   fmt.Println(errBool)
	}
 	inputCatg := category[0]
	inputBind := bindType[0]
	inputPhoto := photoPath[0]
	
	var myBook BookItem
	err := json.NewDecoder(r.Body).Decode(&myBook) 
	if err != nil {
		fmt.Print(err)
	}

	//build the array
	myBook.Id = inputId
	myBook.Name = inputName
	myBook.Author = inputAuthor
	myBook.AuthorEmail = inputAuthorEmail
	myBook.Published = inputPublished
	myBook.Pages = inputPages
	myBook.Publisher = inputPublisher
	myBook.IsAvailable = inputIsAvail
	myBook.Category = inputCatg
	myBook.BindType = inputBind
	myBook.PhotoPath = inputPhoto
    fmt.Println(myBook)
	insertResult, err := bookCollection.InsertOne(context.TODO(), myBook )
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertResult)
	// InsertOne returns a boolean acknowledged as true if the operation ran with write concern 
	// or false if write concern was disabled.
	// A field insertedID with the _id value of the inserted document.
	json.NewEncoder(w).Encode(insertResult.InsertedID)

}

// Get a document based on book id

func getBook(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	ids, ok := r.URL.Query()["id"]
	
	if !ok || len(ids[0]) < 1 {
        log.Println("Url Param 'id' is missing")
        return
	}
	inputId, errConv := strconv.ParseInt(ids[0], 10, 64)
	if errConv != nil {
		fmt.Println("in controllers => getBook 003")
		fmt.Println(errConv)
	}
	
	var myBook BookItem
	e := json.NewDecoder(r.Body).Decode(&myBook)
	if e != nil {
		fmt.Println("in controllers => getBook 001")
		fmt.Println(e)
	}
	fmt.Println(inputId)
	//var result primitive.M //  an unordered representation of a BSON document which is a Map
	result := BookItem{}
	err := bookCollection.FindOne(context.TODO(), bson.D{{"id", inputId}}).Decode(&result)
	//err := bookCollection.FindOne(context.TODO(), bson.D{{"id", 1}}).Decode(&result)
	if err != nil {
		fmt.Println("in controllers => getBook 002")
		fmt.Println(err)
	}
	fmt.Println(result.Id)
	fmt.Println(result.Name)
	fmt.Println(result.Published)
	fmt.Println(result.IsAvailable)
	fmt.Println(result.PhotoPath)
	json.NewEncoder(w).Encode(result) // returns a Map containing document

}

//Update a book

func updateBook(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	// get the values from URL parameters
	ids, ok := r.URL.Query()["id"]
	if !ok || len(ids[0]) < 1 {
        log.Println("Url Param 'id' is missing!")
        return
	}

	name, ok_name := r.URL.Query()["name"]
	if !ok_name || len(name[0]) < 1 {
        log.Println("Url Param 'name' is missing!")
        return
	}

	author, ok_author := r.URL.Query()["author"]
	if !ok_author || len(author[0]) < 1 {
        log.Println("Url Param 'author' is missing!")
        return
	}

	authorEmail, ok_authoremail := r.URL.Query()["authoremail"]
	if !ok_authoremail {
        log.Println("Url Param 'author-email' has issue!")
        return
	}

	published, ok_published := r.URL.Query()["published"]
	if !ok_published {
        log.Println("Url Param 'published' has issue!")
        return
	}

	pages, ok_pages := r.URL.Query()["pages"]
	if !ok_pages {
        log.Println("Url Param 'pages' has issue!")
        return
	}

	publisher, ok_pub := r.URL.Query()["publisher"]
	if !ok_pub {
        log.Println("Url Param 'publisher' has issue!")
        return
	}

	isAvailable, ok_avail := r.URL.Query()["isavailable"]
	if !ok_avail {
        log.Println("Url Param 'isavailable' has issue!")
        return
	}

	category, ok_cat := r.URL.Query()["category"]
	if !ok_cat || len(category[0]) < 1   {
        log.Println("Url Param 'category' is missing!")
        return
	}

	bindType, ok_bind := r.URL.Query()["bindtype"]
	if !ok_bind || len(bindType[0]) < 1   {
        log.Println("Url Param 'bindtype' is missing!")
        return
	}

	photoPath, ok_photo := r.URL.Query()["photopath"]
	if !ok_photo || len(photoPath[0]) < 1   {
        log.Println("Url Param 'photopath' is missing!")
        return
	}
    // get the values from array to local variables
	inputId, errConv := strconv.ParseInt(ids[0], 10, 64)
	if errConv != nil {
		fmt.Println("in controllers => updatebook")
		fmt.Println(errConv)
	}
	inputPages, errConvPages := strconv.ParseInt(pages[0], 10, 64)
	if errConvPages != nil {
		fmt.Println(errConv)
	}
	inputName := name[0]
	inputAuthor := author[0]
	inputAuthorEmail := authorEmail[0]
	inputPublished := published[0]
	inputPublisher := publisher[0]
	inputIsAvail := isAvailable[0]
	inputCatg := category[0]
	inputBind := bindType[0]
	inputPhoto := photoPath[0]

	var myBook BookItem
	e := json.NewDecoder(r.Body).Decode(&myBook)
	if e != nil {
		fmt.Print(e)
	}
	filter := bson.D{{"id", inputId}} // converting value to BSON type
	after := options.After                // for returning updated document
	returnOpt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}
	update := bson.D{{"$set", bson.D{{"name", inputName},{"author",inputAuthor},{"authoremail",inputAuthorEmail},{"published",inputPublished},{"pages",inputPages},{"publisher",inputPublisher},{"isavailable",inputIsAvail},{"category",inputCatg},{"bindtype",inputBind},{"photopath",inputPhoto}}}}
	updateResult := bookCollection.FindOneAndUpdate(context.TODO(), filter, update, &returnOpt)

	var result primitive.M
	_ = updateResult.Decode(&result)

	json.NewEncoder(w).Encode(result)
}

//Delete book

func deleteBook(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	/* Use objectId later, for now deleting document based on id
	params := mux.Vars(r)["_id"] //get Parameter value as string

	_id, err := primitive.ObjectIDFromHex(params) // convert params to mongodb Hex ID
	if err != nil {
		fmt.Printf(err.Error())
	}
	opts := options.Delete().SetCollation(&options.Collation{}) // to specify language-specific rules for string comparison, such as rules for lettercase
	res, err := bookCollection.DeleteOne(context.TODO(), bson.D{{"_id", _id}}, opts)
	*/
	ids, ok := r.URL.Query()["id"]
	if !ok || len(ids[0]) < 1 {
        log.Println("Url Param 'id' is missing!")
        return
	}
	inputId, errConv := strconv.ParseInt(ids[0], 10, 64)
	if errConv != nil {
		fmt.Println("in controllers => updatebook")
		fmt.Println(errConv)
	}

	opts := options.Delete().SetCollation(&options.Collation{}) // to specify language-specific rules for string comparison, such as rules for lettercase
	res, err := bookCollection.DeleteOne(context.TODO(), bson.D{{"id", inputId}}, opts)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("deleted %v documents\n", res.DeletedCount)
	json.NewEncoder(w).Encode(res.DeletedCount) // return number of documents deleted

}

func getAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var results []primitive.M                                   //slice for multiple documents
	cursor, err := bookCollection.Find(context.TODO(), bson.D{{}}) //returns a *mongo.Cursor
	if err != nil {

		fmt.Println(err)

	} 
	for cursor.Next(context.TODO()) { //Next() gets the next document for corresponding cursor

		var elem primitive.M
		err := cursor.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, elem) // appending document pointed by Next()
	}
	cursor.Close(context.TODO()) // close the cursor once stream of documents has exhausted
	json.NewEncoder(w).Encode(results)
}