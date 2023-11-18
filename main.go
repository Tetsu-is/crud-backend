package main

import (
	"crud-backend/middleware"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Todo struct {
	ID   uint   `json:"id"`
	Task string `json:"task"`
}

var db *gorm.DB
var err error

func main() {

	//connect to MySQL database
	dsn := "tetsuro:11quin17@tcp(localhost:3306)/todo?charset=utf8&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect to MySQL database")
		return
	}

	//Automigrate the Todo model
	db.AutoMigrate(&Todo{})

	//Initialize GIN router
	router := gin.Default()

	//cors
	router.Use(middleware.Cors())

	//Define routes
	router.GET("/todos", getTodos)
	router.POST("/todos", addTodo)

	//Start the server
	router.Run("localhost:9090")
}

// Handler function to get all Todos
func getTodos(c *gin.Context) {
	var todos []Todo
	db.Find(&todos)
	c.JSON(200, todos)
}

// Handler function to add a new Todo
func addTodo(c *gin.Context) {
	var todo Todo
	c.BindJSON(&todo)
	db.Create(&todo)
	c.JSON(200, todo)
}
