package controllers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/api/iterator"
	"test.com/config"
	"test.com/models"
)

var collectionName ="todos"

func GetAllTodos(c *gin.Context) {
	iter :=  config.FireStoreClient.Collection(collectionName).Documents(context.Background())
	var todos [] models.Todo
	var todo models.Todo
	for {
			doc, err := iter.Next()
			if err == iterator.Done {
					break
			}
			if err != nil {
					c.AbortWithStatus(500)
			}
			
			doc.DataTo(&todo)
			todos=append(todos, todo)
	}
	c.JSON(http.StatusOK, todos)
}

func GetTodoById(c *gin.Context){
	id:= c.Param("id")
	var todo models.Todo
	doc,err:=  config.FireStoreClient.Collection(collectionName).Doc(id).Get(context.Background())
	if err != nil{
		c.AbortWithError(500,err)
	}
	doc.DataTo(&todo)
	c.JSON(http.StatusOK,todo)
}

func InsertTodo(c *gin.Context){
	var newTodo models.Todo

	if err := c.BindJSON(&newTodo); err != nil{
		c.AbortWithError(500,err)
	}

	docRef:=  config.FireStoreClient.Collection(collectionName).NewDoc()
	_,err:=docRef.Create(context.Background(),newTodo)

	if err != nil{
		c.AbortWithError(500,err)
	}

	c.JSON(http.StatusOK,docRef.ID)
}

func UpdateTodo(c *gin.Context){
	var updatedTodo models.Todo
	if err := c.BindJSON(&updatedTodo); err != nil{
		c.AbortWithError(500,err)
	}
	
	_,err:=config.FireStoreClient.Collection(collectionName).Doc(c.Param("id")).Update(context.Background(),
	updatedTodo.ToFirestoreUpdate(),
	)
	if err != nil{
		c.AbortWithError(http.StatusBadGateway,err)
	}

	c.JSON(http.StatusOK,"Updated")
}

func DeleteTodo(c *gin.Context){
	_,err := config.FireStoreClient.Collection(collectionName).Doc(c.Param("id")).Delete(context.Background())

	if err != nil{
		c.AbortWithError(http.StatusBadGateway,err)
	}

	c.JSON(http.StatusOK,"Deleted")
}