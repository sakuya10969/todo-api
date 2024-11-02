// routes/router.go
package routes

import (
    "todo-api/controllers"
    "github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
    api := router.Group("/api")
    {
        api.GET("/todos", controllers.GetTodos)
        api.GET("/todos/:id", controllers.GetTodoByID)
        api.POST("/todos", controllers.CreateTodo)
        api.PUT("/todos/:id", controllers.UpdateTodo)
        api.DELETE("/todos/:id", controllers.DeleteTodoByID)
    }
}
