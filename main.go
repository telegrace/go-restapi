package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)
type Article struct {
	Title string `json:"Title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
	// := initialises the variable and typeguards too 
	articles:= Articles{
		Article{Title:"Test Title", Desc: "Test Description", Content: "Test Content"},
	}
	fmt.Println("Endpoint Hit: All Articles Endpoint")
	json.NewEncoder(w).Encode(articles) //encoding?
}
//r request
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HomePage Endpoint Hit!")
	fmt.Println("Endpoint Hit: homePage")
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Endpoint Hit!")
	fmt.Println("Endpoint Hit: hello")
}

func handleRequests(){
	myRouter := mux.NewRouter().StrictSlash(true) 
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/hello", hello)
	myRouter.HandleFunc("/articles", allArticles)
	log.Fatal(http.ListenAndServe(":8081", myRouter)) //why is myRouter not nil?
}
func main(){
	handleRequests()
}