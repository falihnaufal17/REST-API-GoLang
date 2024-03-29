package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/users", returnAllUsers).Methods("GET")
	router.HandleFunc("/users", insertUser).Methods("POST")
	http.Handle("/", router)
	fmt.Println("Connected to port 1700")
	log.Fatal(http.ListenAndServe(":1700", router))
}
