package routes

import (
	"github.com/gin-gonic/gin"
	"test.com/controllers"
)

func AddTodoRoutes(route *gin.RouterGroup){
	route.Group("/todo")
	{
		route.GET("getAllTodos",controllers.GetAllTodos)
		route.GET("getTodoById/:id",controllers.GetTodoById)
		route.POST("insertTodo",controllers.InsertTodo)
		route.PUT("updateTodo/:id",controllers.UpdateTodo)
		route.DELETE("deleteTodo/:id",controllers.DeleteTodo)
	}
}