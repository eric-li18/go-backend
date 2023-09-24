package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection
var ctx = context.TODO()

func addHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	location := query.Get("location")
	w.Write([]byte(fmt.Sprintf(`{"location": "%s" }`, location)))
}

func catchAllHandler() http.Handler {
	/* We use an anon function in order to support chaining middleware as
	* a code pattern:
	*
	* 	middleware1(middleware2(finalHandler))
	*
	 */
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Endpoint DNE.")
	})
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/add", addHandler)
	r.PathPrefix("/").Handler(catchAllHandler())
	log.Fatal(http.ListenAndServe(":3000", r))
	fmt.Printf("Server listening on port 3000")
}

func inrit() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database("testDb").Collection("test")
}
