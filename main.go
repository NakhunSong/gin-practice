package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
	Time   int8   `json:"time"`
}

var todos = []todo{
	{ID: "1", Title: "Laundry", Detail: "Single T-shirts", Time: 15},
	{ID: "2", Title: "Shopping", Detail: "Some T-shirts, Pants.", Time: 16},
	{ID: "3", Title: "Dinner", Detail: "Having Dinner at Seoul.", Time: 19},
}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodoByID)
	router.POST("/todos", postTodos)

	router.Run("localhost:8080")
}

func getTodos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, todos)
}

func getTodoByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range todos {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "It was not found"})
}

func postTodos(c *gin.Context) {
	var newTodo todo

	if err := c.BindJSON(&newTodo); err != nil {
		return
	}

	todos = append(todos, newTodo)
	c.IndentedJSON(http.StatusCreated, newTodo)
}
