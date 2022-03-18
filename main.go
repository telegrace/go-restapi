package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"

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
	http.HandleFunc("/", homePage)
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/articles", allArticles)

	log.Fatal(http.ListenAndServe(":8081", nil))
}
func main(){
	handleRequests()
}