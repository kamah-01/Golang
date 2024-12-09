package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
)

type Jobs struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Priority string `json:"priority"`
	Status   string `json:"status"`
	Worker   string `json:"worker"`
}

var jobs = []Jobs{
	{ID: "1", Title: "Kiongozi 1", Priority: "3", Status: "In Progress", Worker: "Mureisi"},
	{ID: "2", Title: "Kiongozi 2", Priority: "2", Status: "To Do", Worker: "Njabe"},
	{ID: "3", Title: "Kiongozi 3", Priority: "2", Status: "In Progress", Worker: "Mzee"},
}

func InsertJob(conn *pgx.Conn, job Jobs) error {
	query := `INSERT INTO jobs (title, priority, status, worker) VALUES ($1, $2, $3, $4)`
	_, err := conn.Exec(context.Background(), query, job.Title, job.Priority, job.Status, job.Worker)
	if err != nil {
		return fmt.Errorf("failed to insert job: %v", err)
	}
	fmt.Println("Job inserted successfully!")
	return nil
}
