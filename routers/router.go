package routers

import "github.com/gin-gonic/gin"

func StartApp() {
	router := gin.Default()

	todosRouter := router.Group("/todos")
	{
		todosRouter.GET("/", controllers.GetTodos)
		todosRouter.POST("/", controllers.CreateTodo)
		todosRouter.GET("/:id", controllers.GetTodo)
		todosRouter.PUT("/:id", controllers.UpdateTodo)
		todosRouter.DELETE("/:id", controllers.DeleteTodo)
	}

	return router
}
