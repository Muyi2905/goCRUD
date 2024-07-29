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

type movies struct{
	id string
}