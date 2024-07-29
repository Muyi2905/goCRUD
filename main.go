package main

import (
	"encoded/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct{
	Id string `json:"id"`
	Title string `json:"title"`
	Director string `json:"director"`
}

type Director struct {
	firstName string `json:"firstname"`
	lastName string `json:"lastname"`
}

var movies []Movie

func main(){
	r:= mux.NewRouter()

	r.HandleFunc("/movies", getmovie).Methods("GET")
	r.HandleFunc("/movie{id}", getmovie).Methods("GET")
	r.HandleFunc("/movie", createmovie).Methods("POST")
	r.HandleFunc("/movies{id}", updatemovie).Methods("PATCH")
	r.HandleFunc("/movies{id}",deletemovie).Methods("DELETE")

	fmt.Printf("server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}