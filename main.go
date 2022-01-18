package main

import (
	"log"
	"net/http"

	"github.com/NeerChayaphon/Task-Management-API-with-Go/database"
	"github.com/gin-gonic/gin"
)

func main() {

	// Database setup
	db, err := database.NewDatabase()
	if err != nil {
		log.Fatal(err)
	}
	database.MigrateDB(db)
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
