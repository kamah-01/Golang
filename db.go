package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
)

func ConnectDB(ctx context.Context) (*pgx.Conn, error) {
	const connStr = "postgresql://kamah:uhDlDVCjFMQSpcpEHhNK6g@blogkamah-4995.j77.aws-ap-south-1.cockroachlabs.cloud:26257/gin?sslmode=verify-full"

	connConfig, err := pgx.ParseConfig(connStr)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}

	conn, err := pgx.ConnectConfig(ctx, connConfig)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}

	err = checkConnection(conn)
	if err != nil {
		conn.Close(ctx)
		return nil, fmt.Errorf("database connection check failed: %v", err)
	}

	err = CreateTable(conn)
	if err != nil {
		conn.Close(ctx)
		return nil, fmt.Errorf("failed to create jobs table: %v", err)
	}

	return conn, nil
}

func checkConnection(conn *pgx.Conn) error {

	var dummy int
	err := conn.QueryRow(context.Background(), "SELECT 1").Scan(&dummy)
	if err != nil {
		return fmt.Errorf("failed to run test query: %v", err)
	}
	fmt.Println("Database connection successful")
	return nil
}
