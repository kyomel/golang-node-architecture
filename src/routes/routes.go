package routes

import (
	"todolist-golang/src/controllers"

	"github.com/gin-gonic/gin"
)

// Routes function to serve endpoints
func Routes() {
	route := gin.Default()

	route.POST("/todo", controllers.CreateTodo)
	route.GET("/todo", controllers.GetAllTodos)
	route.PUT("/todo/:idTodo", controllers.UpdateTodo)
	route.DELETE("/todo/:idTodo", controllers.DeleteTodo)

	route.Run()
}
