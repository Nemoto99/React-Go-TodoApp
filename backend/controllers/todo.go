package controllers

import (
	"app/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetTodos(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var todos []models.Todo
		result := db.Find(&todos)

		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Can not read todos"})
			return
		}
		c.JSON(http.StatusOK, &todos)
	}
}

func PostTodo(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newTodo models.Todo

		if err := c.BindJSON(&newTodo); err != nil {
			return
		}

		result := db.Create(&newTodo)

		if result.RowsAffected == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Can not create todo"})
			return
		}

		c.IndentedJSON(http.StatusCreated, newTodo)
	}
}

func UpdateTodo(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		var updateTodo models.Todo
		if err := c.ShouldBind(&updateTodo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to bind data"})
			return
		}

		var todo models.Todo
		if result := db.First(&todo, id); result.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
			return
		}

		todo.Completed = updateTodo.Completed

		if result := db.Save(&todo); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update todo"})
			return
		}

		c.JSON(http.StatusOK, todo)
	}
}

func DeleteTodo(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		result := db.Where("id = ?", id).Delete(&models.Todo{})

		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Can not delete todo"})
			return
		}

		if result.RowsAffected == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Not found the todo"})
			return
		}
	}
}
