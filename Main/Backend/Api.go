package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type Pos struct {
}

var Positions []Pos

func GameInput(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

	}
}

func GameGet(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {

	}
}

func GameDelete(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {

	}
}

func main() {
	router := chi.NewRouter()

	http.HandleFunc("/Game/Post", GameInput)
	http.HandleFunc("Game/Get", GameGet)
	http.HandleFunc("/Game/Delete", GameDelete)

	fmt.Print("Server is running")
	http.ListenAndServe("localhost:25025", router)
}
