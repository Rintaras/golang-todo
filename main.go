package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
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

	//URLからIDのみを抽出する
	idStr := strings.TrimPrefix(r.URL.Path, "/todos/")

	if idStr == "" {
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
	return
	}

	id, err := strconv.Atoi(idStr)

	//無効なやつの処理
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte(`{"error": "無効なID形式です"}`))
        return
    }


	//有効IDに対する処理
	switch r.Method {

	// GET /todos/{id} のTodoを返す
    case http.MethodGet:
        for _, t := range todos {
            if t.ID == id {
                json.NewEncoder(w).Encode(t)
                return
            }
        }
        w.WriteHeader(http.StatusNotFound)
        w.Write([]byte(`{"error": "Todoが見つかりません"}`))

	// DELETE /todos/{id} のTodoを削除する
	case http.MethodDelete:
	for index, t := range todos {
		if t.ID == id {
			todos = append(todos[:index], todos[index+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"error": "削除するTodoが見つかりません"}`))

    default:
        w.WriteHeader(http.StatusMethodNotAllowed)
    }

}

func main() {
	http.HandleFunc("/todos/", getTodos)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("サーバーの起動に失敗 : ", err)
	}

}

