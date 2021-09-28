package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)
const (
    ContentTypeBinary = "application/octet-stream"
    ContentTypeForm   = "application/x-www-form-urlencoded"
    ContentTypeJSON   = "application/json"
    ContentTypeHTML   = "text/html; charset=utf-8"
    ContentTypeText   = "text/plain; charset=utf-8"
)

func SetupRouter() *gin.Engine {
	routeEngine := gin.Default()
	routeEngine.GET("/", func(c *gin.Context) {
        c.Data(http.StatusOK,"text/html; charset=utf-8",[]byte("<div><h1>Welcome to TODO API</h1><h2>Everything is fine...</h2></div>"))
    })
	// routeEngine.Use(middleware.CORSMiddleware())
	// routeEngine.Use(middleware.AuthMiddleware)
	todoRoute := routeEngine.Group("/todo")
	AddTodoRoutes(todoRoute)
	return routeEngine
}