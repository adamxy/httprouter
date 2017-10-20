package main

import (
	"log"
	"net/http"
    "github.com/adamxy/httprouter/router"
)

func main() {
	router := NewRouter(AllRoutes())
	log.Fatal(http.ListenAndServe(":8080", router))
}
