package controllers

import (
	"fmt"
	"net/http"

	"todolist-golang/src/config"

	"todolist-golang/src/models"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"gorm.io/gorm"
)

// Define database client
var db *gorm.DB = config.ConnectDB()

// Todo struct for request body
type todoRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Defining struct for response
type todoResponse struct {
	todoRequest
	ID uint `json:"id"`
}

// Create todo data to database by run this function
func CreateTodo(context *gin.Context) {
	var data todoRequest
	// Binding request body json to request body struct
	if err := context.ShouldBindJSON(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Matching todo models struct with todo request struct
	todo := models.Todo{}
	todo.Name = data.Name
	todo.Description = data.Description

	// Querying to database
	result := db.Create(&todo)
	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong"})
		return
	}

	// Matching result to create response
	var response todoResponse
	response.ID = todo.ID
	response.Name = todo.Name
	response.Description = todo.Description

	// Creating http response
	context.JSON(http.StatusCreated, response)
}

// Getting all todo datas
func GetAllTodos(context *gin.Context) {
	var todos []models.Todo

	// Querying to find todo datas
	err := db.Find(&todos)
	if err.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Error geeting data"})
		return
	}

	// Creating http response
	context.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Success",
		"data":    todos,
	})
}

// Update todo data
func UpdateTodo(context *gin.Context) {
	var data todoRequest

	// Defining request parameter to get todo id
	reqParamId := context.Param("idTodo")
	idTodo := cast.ToUint(reqParamId)

	// Binding request body json to request body struct
	if err := context.BindJSON(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Initiate models todo
	todo := models.Todo{}

	// Querying find todo data by todo id from request parameter
	todoById := db.Where("id = ?", idTodo).First(&todo)
	if todoById.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Todo not found"})
		return
	}

	// Matching todo request with todo models
	todo.Name = data.Name
	todo.Description = data.Description

	// Update new todo data
	result := db.Save(&todo)
	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong"})
		return
	}

	// Matching result to todo response struct
	var response todoResponse
	response.ID = todo.ID
	response.Name = todo.Name
	response.Description = todo.Description

	// Creating http response
	context.JSON(http.StatusCreated, response)
}

// Delete todo data function
func DeleteTodo(context *gin.Context) {
	// Initiate todo models
	todo := models.Todo{}
	// Getting request parameter id
	reqParamId := context.Param("idTodo")
	idTodo := cast.ToUint(reqParamId)

	// Querying delete todo by id
	delete := db.Where("id = ?", idTodo).Unscoped().Delete(&todo)
	fmt.Println(delete)

	// Creating http response
	context.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Success",
		"data":    idTodo,
	})

}
