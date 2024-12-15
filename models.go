package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
)

func CreateTable(conn *pgx.Conn) error {
	query := `
	CREATE TABLE IF NOT EXISTS jobs (
		id SERIAL PRIMARY KEY,
		title TEXT NOT NULL,
		priority INT NOT NULL,
		status TEXT NOT NULL,
		worker TEXT NOT NULL
	);`
	_, err := conn.Exec(context.Background(), query)
	if err != nil {
		return fmt.Errorf("failed to create table: %v", err)
	}
	fmt.Println("Jobs table is ready")
	return nil
}
