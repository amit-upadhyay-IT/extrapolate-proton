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
	router.HandleFunc("/dynamic-programming/questions", questionResources.ViewDynamicProgrammingQuestions)
	router.HandleFunc("/tree/questions", questionResources.ViewTreeQuestions)
	router.HandleFunc("/design-pattern/questions", questionResources.ViewDesignPatternQuestions)
	router.HandleFunc("/system-design/questions", questionResources.ViewSystemDesignQuestions)
	router.HandleFunc("/graph/questions", questionResources.ViewGraphQuestions)

	fmt.Println("starting server")
	err := http.ListenAndServe(":8081", router)
	if err != nil {
		log.Fatal(err)
	}
}
