package main

import (
	"crud-backend/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Todo struct {
	ID   uint   `json:"id"`
	Task string `json:"task"`
}

var db *gorm.DB
var err error

func main() {

	// connect to MySQL database
	// dsn := "crud:pARlURjyh8PNe7K8@tcp(localhost:3306)/todo?charset=utf8&parseTime=True&loc=Local"
	// db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	fmt.Println("Failed to connect to MySQL database")
	// 	return
	// }

	// var err error
	// dsn := "host=dpg-clcpb9bmot1c73di45t0-a user=crud_6hq6_user password=u2AXJwBWKyv4hiHVryZjWkPF682L96sS dbname=crud_6hq6 port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	// db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// if err != nil {
	// 	fmt.Println("Failed to connect DB")
	// }

	// //Automigrate the Todo model
	// db.AutoMigrate(&Todo{})

	//Initialize GIN router
	router := gin.Default()

	//cors
	router.Use(middleware.Cors())

	//Define routes
	router.GET("/todos", getTodos)
	router.POST("/todos", addTodo)
	router.GET("/hello", hello)

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

func hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello world",
	})
}
