package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	Id string  `json:"id"`
	Title string `json:"title"`
	Completed bool `json:"completed"`
}

var todoList = []todo{
{
	Id:"1",
	Title:"Learn more about GO",
	Completed:false,
},
{
	Id:"2",
	Title:"Learn more about GIN",
	Completed:false,
},
}

func getTodoList(ctx *gin.Context){
	ctx.IndentedJSON(http.StatusOK,todoList)
}
func addTodo(ctx *gin.Context){
	var newTodo todo
	if err:= ctx.BindJSON(&newTodo);err!=nil{
		return
	}
	todoList = append(todoList, newTodo)
	ctx.IndentedJSON(http.StatusCreated,newTodo)
}
func getTodoByID(id string)(*todo,error) {
	for i,t := range todoList {
		if t.Id == id {
			return &todoList[i],nil
		}
	}
	return nil, errors.New("todo not exist")
}
func getTodo(ctx *gin.Context){
	id:= ctx.Param("id")
	todo,err := getTodoByID(id)
	if err!= nil{
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message":"todo not found"})
		return 
	}
	ctx.IndentedJSON(http.StatusOK,todo)
}
func UpdateTodoByID(ctx *gin.Context){
	id:= ctx.Param("id")
	_,err := getTodoByID(id)
	if err!= nil{
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message":"todo not found"})
		return 
	}
	var updatedTodo todo
	if err := ctx.BindJSON(&updatedTodo); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	for i, todo := range todoList {
		if todo.Id == id {
			// Update the fields provided in the request
			if updatedTodo.Title != "" {
				todoList[i].Title = updatedTodo.Title
			}
			todoList[i].Completed = updatedTodo.Completed
			

			
			

			ctx.IndentedJSON(http.StatusOK, gin.H{"message": "User updated", "user": todoList[i]})
			return
		}
	}
	// If user not found
	ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})


}

func main() {
	router:= gin.Default()
	router.GET("/todo",getTodoList)
	router.GET("/todo/:id",getTodo)
	router.PATCH("/todo/:id",getTodo)
	router.POST("/todo",addTodo)
	router.Run("localhost:9090")
}