package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func NewJobHandler(c *gin.Context) {

	db, ok := c.MustGet("db").(*pgx.Conn)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection error"})
		return

	}

	var job Jobs
	if err := c.ShouldBindJSON(&job); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	err := InsertJob(db, job)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"message": "Job created successfully"})

}

func GetJobsHandler(c *gin.Context) {
	c.JSON(http.StatusOK, jobs)
}

func UpdateJobHandler(c *gin.Context) {
	id := c.Param("id")
	var job Jobs
	if err := c.ShouldBindJSON(&job); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	index := -1
	for i := 0; i < len(jobs); i++ {
		if jobs[i].ID == id {
			index = i
		}
	}
	if index == -1 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Job not Found",
		})
		return
	}
	jobs[index] = job
	c.JSON(http.StatusOK, jobs)
}

func DeleteJobHandler(c *gin.Context) {
	id := c.Param("id")

	index := -1
	for i := 0; i < len(jobs); i++ {
		if jobs[i].ID == id {
			index = i
		}
	}
	if index == -1 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Job not found",
		})
		return
	}
	jobs = append(jobs[:index], jobs[index+1:]...)
	c.JSON(http.StatusOK, gin.H{
		"error": "Job has been deleted",
	})
}
