package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Todo struct {
    ID      int    `json:"id"`
    Title   string `json:"title"`
    Checked bool   `json:"checked"`
}

var todos = []Todo{
    {ID: 1, Title: "Buy groceries", Checked: false},
    {ID: 2, Title: "Walk the dog", Checked: true},
}

func main() {
    r := gin.Default()

    r.GET("/todos", func(c *gin.Context) {
        c.JSON(http.StatusOK, todos)
    })

    // POST /todos 
    r.POST("/todos", func(c *gin.Context) {
        var newTodo Todo
        if err := c.ShouldBindJSON(&newTodo); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "不正なデータです"})
            return
        }
        todos = append(todos, newTodo)
        c.JSON(http.StatusCreated, newTodo)
    })

    // GET /todos/:id 
    r.GET("/todos/:id", func(c *gin.Context) {

        id, err := strconv.Atoi(c.Param("id"))
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "無効なID形式です"})
            return
        }

        for _, t := range todos {
            if t.ID == id {
                c.JSON(http.StatusOK, t)
                return
            }
        }
        c.JSON(http.StatusNotFound, gin.H{"error": "Todoが見つかりません"})
    })

    // PUT /todos/:id
    r.PUT("/todos/:id", func(c *gin.Context) {
        id, err := strconv.Atoi(c.Param("id"))
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "無効なID形式です"})
            return
        }

        var updatedTodo Todo
        if err := c.ShouldBindJSON(&updatedTodo); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "不正なデータです"})
            return
        }

        for index, t := range todos {
            if t.ID == id {
                todos[index] = updatedTodo
                c.JSON(http.StatusOK, updatedTodo)
                return
            }
        }
        c.JSON(http.StatusNotFound, gin.H{"error": "更新するTodoが見つかりません"})
    })

    // DELETE /todos/:id 
    r.DELETE("/todos/:id", func(c *gin.Context) {
        id, err := strconv.Atoi(c.Param("id"))
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "無効なID形式です"})
            return
        }

        for index, t := range todos {
            if t.ID == id {
                todos = append(todos[:index], todos[index+1:]...)
                c.Status(http.StatusNoContent)
                return
            }
        }
        c.JSON(http.StatusNotFound, gin.H{"error": "削除するTodoが見つかりません"})
    })
    r.Run()
}