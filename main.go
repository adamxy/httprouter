package main

import (
	"log"
	"net/http"
    "github.com/adamxy/apigo/httprouter"
)

func main() {
	router := NewRouter(AllRoutes())
	log.Fatal(http.ListenAndServe(":8080", router))
}
