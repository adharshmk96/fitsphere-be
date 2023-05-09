package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

// Create a global connection pool variable
var dbPool *pgxpool.Pool

// Connect establishes a connection to the PostgreSQL database.
func Connect(host, port, user, password, database string) error {
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, database)

	// Connect to the database using connectionString
	poolConfig, err := pgxpool.ParseConfig(connectionString)
	if err != nil {
		return err
	}

	dbPool, err = pgxpool.ConnectConfig(context.Background(), poolConfig)
	if err != nil {
		return err
	}

	return nil
}

// GetDBPool returns the current database connection pool for use in other packages.
func GetDBPool() *pgxpool.Pool {
	return dbPool
}

// Close terminates the connection pool.
func Close() {
	dbPool.Close()
}
