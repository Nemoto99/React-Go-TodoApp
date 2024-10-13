package main

import (
	"app/controllers"
	"app/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	r := gin.Default()
	db := ConnectDatabase()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	r.GET("/todos", controllers.GetTodos(db))
	r.POST("/todo", controllers.PostTodo(db))
	// r.PUT("/todo/:id", controllers.UpdateTodo(db))
	r.DELETE("/todo/:id", controllers.DeleteTodo(db))

	r.Run(":8080")
}

func ConnectDatabase() *gorm.DB {
	dsn := "docker:docker@tcp(database)/main?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect db")
	}

	db.AutoMigrate(&models.Todo{})

	return db
}
