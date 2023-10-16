package main

// get gin
// dependecy tracking go mod init  go mod init example/todo-go
// download gin
//go get github.com/gin-gonic/gin

// api

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// json eq name
type todo struct {
	// name type
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

// not json
var todos = []todo{
	{ID: "1", Item: "YOLO", Completed: false},
	{ID: "2", Item: "YOLOsadf", Completed: false},
	{ID: "3", Item: "YOLsdfO", Completed: false},
}

// client and server will comm in json
// server need to convert this to json vice versa client
// get net/http and github.com/gin

func getTodos(ctx *gin.Context) {
	// bunch info of http request
	// conver to json
	ctx.IndentedJSON(http.StatusOK, todos)
}

func addTodo(ctx *gin.Context) {
	// we need to infor from request
	var newTodo todo
	if err := ctx.BindJSON(&newTodo); err != nil {
		// bind json from request to newToDo
		return
	}
	todos = append(todos, newTodo)
	ctx.IndentedJSON(http.StatusCreated, todos)
}

func getTodo(ctx *gin.Context) {
	// extract path paramerter from url
	id := ctx.Param("id")
	todo, err := getTodoByID(id)
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "no found"})
		return
	}

	ctx.IndentedJSON(http.StatusOK, todo)
}

func getTodoByID(id string) (*todo, error) {
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}
	return nil, errors.New("to do not found")
}

func mainOld() {
	// to create server create router
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.POST("/todos", addTodo)

	router.GET("/todos/:id", getTodo)
	// run by go run main.go
	router.Run("localhost:5001")
}
