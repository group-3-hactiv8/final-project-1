package controllers

import (
	"final-project-1/database"
	"final-project-1/dto"
	"final-project-1/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTodos(c *gin.Context) {
	db := database.GetDB()
	var todos []models.Todo

	if err := db.Debug().Model(&models.Todo{}).Find(&todos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	var todosResponse []dto.TodoResponse

	for _, todo := range todos {
		todoResponse := dto.TodoResponse{
			ID:        todo.ID,
			UserID:    todo.UserID,
			Title:     todo.Title,
			Completed: todo.Completed,
		}
		todosResponse = append(todosResponse, todoResponse)
	}

	c.JSON(http.StatusOK, todosResponse)
}

func CreateTodo(c *gin.Context) {
	var newTodo models.Todo
	if err := c.ShouldBindJSON(&newTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	db := database.GetDB()

	var exists bool
	err := db.Debug().Model(&models.Todo{}).
		Select("count(*) > 0").
		Where("title = ?", newTodo.Title).
		Find(&exists).
		Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	if exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Todo already exists",
		})
		return
	}

	if err := db.Debug().Create(&newTodo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	TodoResponse := dto.TodoResponse{
		ID:        newTodo.ID,
		UserID:    newTodo.UserID,
		Title:     newTodo.Title,
		Completed: newTodo.Completed,
	}

	c.JSON(http.StatusCreated, TodoResponse)
}
