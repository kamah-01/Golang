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
		title STRING NOT NULL,
		priority STRING NOT NULL,
		status STRING NOT NULL,
		worker STRING NOT NULL
	);`
	_, err := conn.Exec(context.Background(), query)
	if err != nil {
		return fmt.Errorf("failed to create table: %v", err)
	}
	fmt.Println("Jobs table is ready")
	return nil
}
