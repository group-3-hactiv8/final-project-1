package routers

import (
	"final-project-1/controllers"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	router := gin.Default()

	todosRouter := router.Group("/todos")
	{
		todosRouter.GET("/", controllers.GetTodos)
		todosRouter.POST("/", controllers.CreateTodo)
		// todosRouter.GET("/:id", controllers.GetTodo)
		// todosRouter.PUT("/:id", controllers.UpdateTodo)
		// todosRouter.DELETE("/:id", controllers.DeleteTodo)
	}

	return router
}
