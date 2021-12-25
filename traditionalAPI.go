package main

//TODO: Está é uma REST API tradicional utilizando net/http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	Articles = []Article{
		Article{Title: "Olá", Desc: "Article Description", Content: "Article Content"},
		Article{Title: "Olá denovo", Desc: "Article Description", Content: "Article Content"},
	}
	handleRequests()
}
func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", returnAllArticles)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bem vindo a minha Página para API")
	fmt.Println("Endpoint Hit: PaginaInicial")
}

type Article struct {
	Title   string `json:"Título"`
	Desc    string `json:"Descrição"`
	Content string `json:"Conteúdo"`
}

var Articles []Article

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}
