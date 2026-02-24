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


var todos = []Todo{
		{ID: 1, Title: "Buy groceries", Checked: false},
		{ID: 2, Title: "Walk the dog", Checked: true},
	}

func getTodos(w http.ResponseWriter, r *http.Request) {
	
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
	json.NewEncoder(w).Encode(todos)

    case http.MethodPost:
		var newTodo Todo
		if err := json.NewDecoder(r.Body).Decode(&newTodo); err != nil {
            w.WriteHeader(http.StatusBadRequest)
            return
        }
        todos = append(todos, newTodo)
        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(newTodo)

    default:
        w.WriteHeader(http.StatusMethodNotAllowed)
    }

}

func main() {
	http.HandleFunc("/todos", getTodos)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("サーバーの起動に失敗 : ", err)
	}

}

