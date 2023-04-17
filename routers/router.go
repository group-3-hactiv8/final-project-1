package routers

import (
	"final-project-1/controllers"

	"github.com/gin-gonic/gin"

	_ "final-project-1/docs"

	ginSwagger "github.com/swaggo/gin-swagger"

	swaggerfiles "github.com/swaggo/files"
)

// @title Todos API
// @version 1.0
// @description This is a sample server for managing Todo.
// @termsOfService http://swagger.io/terms/
// @contact.name Swagger API Team
// @host localhost:8080
// @BasePath /
func StartApp() *gin.Engine {
	router := gin.Default()

	todosRouter := router.Group("/todos")
	{
		todosRouter.GET("/", controllers.GetTodos)
		todosRouter.POST("/", controllers.CreateTodo)
		// todosRouter.GET("/:id", controllers.GetTodo)
		todosRouter.PUT("/:id", controllers.UpdateTodo)
		// todosRouter.DELETE("/:id", controllers.DeleteTodo)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return router
}
