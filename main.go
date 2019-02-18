package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

func main() {

	port := ":80"

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", home).Methods("GET")

	log.Fatal(http.ListenAndServe(port, router))

	fmt.Println("Server running in port:", port)
}

func home(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/index.html"))
	if err := t.ExecuteTemplate(w, "index.html", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
