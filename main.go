package main

import (
	"fmt"
	"net/http"
	"encoded/json"
	"math/rand"
	"log"
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
	
}