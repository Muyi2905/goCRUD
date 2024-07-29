package main

import (
	"fmt"
	"log"
	"net/http"
	"encoded/json"
	"strconv"
	"math/rand"
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

func main(){
movies = append(movies, Movie{Id: "1", Title: "Interstallar", Director: &Director{firstName:"Christopher", lastName: "Nolan"}})
movies= append(movies, Movie{Id: "2", Title: "Dune2", Director: &Director{firstName: "Denis", lastName: "Villenuve"} })


}