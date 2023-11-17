package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"title"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{ID: "1", Item: "部屋掃除", Completed: false},
	{ID: "2", Item: "本を読む", Completed: false},
	{ID: "3", Item: "買い物", Completed: false},
}

// jsonに変換
func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}
func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.Run("localhost:9090")
}
