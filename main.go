package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Todo struct {
	ID   int `json:"id"`
	Title string `json:"title"`
	Checked bool `json:"checked"`
}

func getTodos(w http.ResponseWriter, r *http.Request) {
	
	w.Header().Set("Content-Type", "application/json")

	todos := []Todo{
		{ID: 1, Title: "Buy groceries", Checked: false},
		{ID: 2, Title: "Walk the dog", Checked: true},
	}

	json.NewEncoder(w).Encode(todos)
}

func main() {
	http.HandleFunc("/todos", getTodos)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("サーバーの起動に失敗 : ", err)
	}

}

