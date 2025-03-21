package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	_ "github.com/go-chi/chi/v5"
	"net/http"
)

type Pos struct {
	Id    int    `json:"id"`
	Value string `json:"value"`
}

var Positions []Pos

// Kolko - 1 || krzyzyk - 2
var Switch int

func CreateBoard() {
	for i := 0; i < 9; i++ {
		Positions = append(Positions, Pos{i, ""})
	}
	Switch = 1
}

func GameInput(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var holder Pos
		json.NewDecoder(r.Body).Decode(&holder)

		if Positions[holder.Id].Value == "" {
			switch Switch {
			case 1:
				Positions[holder.Id].Value = "O"
			case 2:
				Positions[holder.Id].Value = "X"
			}
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Add("Action", "Success")
	}
}

func GameGet(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Header().Add("content-type", "application/json")
		w.WriteHeader(http.StatusOK)

		holder := Positions

		w.WriteHeader(http.StatusOK)
		w.Header().Add("Action", "Success")

		json.NewEncoder(w).Encode(holder)
	}
}

func GameDelete(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		for i := 0; i < 9; i++ {
			Positions[i].Value = ""
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Add("Action", "Success")
	}
}

func main() {
	router := chi.NewRouter()
	CreateBoard()

	http.HandleFunc("/Game/Post", GameInput)
	http.HandleFunc("Game/Get", GameGet)
	http.HandleFunc("/Game/Delete", GameDelete)

	fmt.Print("Server is running")
	http.ListenAndServe("18.156.42.200:25025", router)
}
