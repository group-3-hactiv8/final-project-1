package controllers

import (
	"final-project-1/database"
	"final-project-1/dto"
	"final-project-1/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllTodos godoc
// @Tags todos
// @Description Get All Todos Data
// @ID get-all-Todos
// @Produce json
// @Success 200 {object} dto.Response
// @Router /todos [get]
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

	response := dto.Response{
		Message:    "Todos fetched successfully",
		StatusCode: http.StatusOK,
		Data:       todosResponse,
	}

	c.JSON(http.StatusOK, response)
}

// CreateNewTodo godoc
// @Tags todos
// @Description Create New Todo Data
// @ID create-new-Todo
// @Accept json
// @Produce json
// @Param models.Todo body models.Todo true "request body json"
// @Success 201 {object} dto.Response
// @Router /todos [post]
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

	response := dto.Response{
		Message:    "Todo created successfully",
		StatusCode: http.StatusCreated,
		Data:       TodoResponse,
	}

	c.JSON(http.StatusCreated, response)
}

// UpdateATodo godoc
// @Tags todos
// @Description Update Todo Data
// @ID update-a-todo
// @Accept json
// @Produce json
// @Param id path uint true "request id"
// @Param models.Todo body models.Todo true "request body json"
// @Success 200 {object} dto.Response
// @Router /todos/{id} [put]
func UpdateTodo(c *gin.Context) {
	var newTodo models.Todo

	db := database.GetDB()

	if err := db.Where("id = ?", c.Param("id")).First(&newTodo).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Todo Not Found !",
		})
		return
	}

	var todoUpdate dto.TodoUpdate

	if err := c.ShouldBindJSON(&todoUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := db.Model(&newTodo).Update("completed", false).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	if err := db.Model(&newTodo).Updates(models.Todo{
		UserID:    todoUpdate.UserID,
		Title:     todoUpdate.Title,
		Completed: todoUpdate.Completed}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	TodoResponse := dto.TodoResponse{
		ID:        newTodo.ID,
		UserID:    newTodo.UserID,
		Title:     newTodo.Title,
		Completed: newTodo.Completed,
	}

	response := dto.Response{
		Message:    "Todo Update successfully",
		StatusCode: http.StatusOK,
		Data:       TodoResponse,
	}

	c.JSON(http.StatusOK, response)
}
