package main

import (
	"./resources"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", resources.ServerHealthCheck)

	questionResources := resources.GetQuestionResource()
	router.HandleFunc("/array/questions", questionResources.ViewArrayQuestions)
	router.HandleFunc("/recursion/questions", questionResources.ViewRecursionQuestions)

	fmt.Println("starting server")
	err := http.ListenAndServe(":8081", router)
	if err != nil {
		log.Fatal(err)
	}
}
