package main

import (
	"context"

	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Use(cors.Default())

	database, err := ConnectDB(context.Background())
	if err != nil {
		log.Printf("Error: could not connect to db, Message: %s", err.Error())
		return
	}

	router.Use(func(c *gin.Context) {
		c.Set("db", database)
		c.Next()
	})

	SetUpRoutes(router, database)

	err = router.Run(":8080")
	if err != nil {
		log.Printf("Error: server failed to start, Message: %s", err.Error())
	}
}
