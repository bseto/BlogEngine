package main

import (
	//"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting Server")

	//r := mux.NewRouter()
	http.Handle("/", http.FileServer(http.Dir(".")))
	log.Println("Listening on 8080")
	http.ListenAndServe(":8080", nil)

}
