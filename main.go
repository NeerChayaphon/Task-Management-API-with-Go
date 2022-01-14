package main

import (
	"net/http"

	"github.com/NeerChayaphon/Task-Management-API-with-Go/database"
	"github.com/NeerChayaphon/Task-Management-API-with-Go/model"
	"github.com/gin-gonic/gin"
)

func main() {

	db, err := database.NewDatabase()
	if err != nil {
		panic("fail to connect database")
	}

	db.AutoMigrate(&model.Task{})

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong!",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
