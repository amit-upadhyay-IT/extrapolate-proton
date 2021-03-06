package resources

import (
	"../services"
	"fmt"
	"net/http"
)

/*
 * this file is suppose to have the end points for questions on different topics in cs
 * eg: /v1/questions/arrays
 */

type QuestionsResource struct {
	questionService *services.QuestionsService
}

/*
 * Make sure this instance is singleton
 */
func GetQuestionResource() *QuestionsResource {
	return &QuestionsResource{services.GetQuestionsService()}
}

func (questionsRes *QuestionsResource)ViewArrayQuestions(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s",  questionsRes.questionService.ArrayQuestions())

}

func (questionsRes *QuestionsResource)ViewRecursionQuestions(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s",  questionsRes.questionService.RecursionQuestions())
}

func (questionsRes *QuestionsResource)ViewDynamicProgrammingQuestions(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s",  questionsRes.questionService.DynamicProgrammingQuestions())
}


func (questionsRes *QuestionsResource)ViewTreeQuestions(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s",  questionsRes.questionService.TreeQuestions())
}

func (questionsRes *QuestionsResource)ViewDesignPatternQuestions(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s",  questionsRes.questionService.DesignPatternQuestions())
}


func (questionsRes *QuestionsResource)ViewSystemDesignQuestions(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s",  questionsRes.questionService.SystemDesignQuestions())
}

func (questionsRes *QuestionsResource)ViewGraphQuestions(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s",  questionsRes.questionService.GraphQuestions())
}

