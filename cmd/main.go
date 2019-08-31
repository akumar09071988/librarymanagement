package main

import (
	"fmt"
	"librarymangement/pkg/routing"
	"log"
	"net/http"
)

func main() {
	router := new(routing.Router)
	fmt.Println("starting server")
	log.Fatal(http.ListenAndServe(":8080", router.Routes()))
}
