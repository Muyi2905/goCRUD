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
	Director *Director `json:"director"`

}

type Director struct{
	firstName string
	lastName string 
}

var movies[]Movie

func getmovies(w http.ResponseWriter, r*http.Request){
	if err:= r.URL.Path!= "GET"; err!=nil{
http.Error(w, "error", http.StatusNotFound)
	}
}

func main(){
movies = append(movies, Movie{Id: "1", Title: "Interstallar", Director: &Director{firstName:"Christopher", lastName: "Nolan"}})
movies= append(movies, Movie{Id: "2", Title: "Dune2", Director: &Director{firstName: "Denis", lastName: "Villenuve"} })

r:= mux.NewRouter()
r.HandleFunc("/movies",getmovies).Methods("GET")
r.HandleFunc("/movies{id}", getmovie).Methods("GET")
r.HandleFunc("/movies",createmovie).Methods("POST")
r.HandleFunc("/movies{id}", updatemovie).Methods("PATCH")
r.Handle("/movies{id}", deletemovie).Methods("DELETE")

fmt.Println("server is running on port 8080")
log.Fatal(http.ListenAndServe(";8080", r))
}