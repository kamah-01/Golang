package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func HomeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Task mgt api"})
}

func SetUpRoutes(router *gin.Engine, db *pgx.Conn) {

	router.GET("/", HomeHandler)
	router.GET("/jobs", GetJobsHandler)
	router.POST("/job", NewJobHandler)
	router.PUT("/job/:id", UpdateJobHandler)
	router.DELETE("/job/:id", DeleteJobHandler)
}
