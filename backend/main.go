package main

import (
	"app/controllers"
	"app/models"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	log.Println("Starting sever...")

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
	r.POST("/todos", controllers.PostTodo(db))
	r.PUT("/todos/:id", controllers.UpdateTodo(db))
	r.DELETE("/todos/:id", controllers.DeleteTodo(db))

	log.Println("Running server on port 8080...")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
	// r.Run(":8080")

}

func ConnectDatabase() *gorm.DB {
	dsn := "docker:docker@tcp(database:3306)/main?charset=utf8mb4&parseTime=True&loc=Local"
	var db *gorm.DB
	var err error
	for i := 0; i < 5; i++ {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}
		time.Sleep(2 * time.Second)
	}
	if err != nil {
		panic("Failed to connect database!")
	}

	db.AutoMigrate(&models.Todo{})

	return db
}
