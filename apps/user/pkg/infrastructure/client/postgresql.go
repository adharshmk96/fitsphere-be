package client

import (
	"github.com/adharshmk96/fitsphere-be/apps/user/pkg/infrastructure/config"
	"github.com/adharshmk96/fitsphere-be/libs/db"
	"github.com/jackc/pgx/v4/pgxpool"
)

var (
	connection *pgxpool.Pool
)

func init() {
	var err error

	connection, err = connectToPostgres()

	if err != nil {
		panic("error connecting to database")
	}
}

func connectToPostgres() (*pgxpool.Pool, error) {
	err := db.Connect(
		config.Get().PG_HOST,
		config.Get().PG_PORT,
		config.Get().PG_USER,
		config.Get().PG_PASSWORD,
		config.Get().PG_DATABASE,
	)

	if err != nil {
		return nil, err
	}

	connection := db.GetDBPool()

	return connection, nil
}

func GetPostgresConnection() *pgxpool.Pool {
	return connection
}
