package client

import (
	"github.com/adharshmk96/fitsphere-be/apps/user/pkg/infrastructure/config"
	"github.com/adharshmk96/fitsphere-be/libs/stk/db"
	"github.com/jackc/pgx/v4/pgxpool"
)

var (
	connection *pgxpool.Pool
)

func init() {
	var err error

	host := config.Get().PG_HOST
	port := config.Get().PG_PORT
	user := config.Get().PG_USER
	password := config.Get().PG_PASSWORD
	database := config.Get().PG_DATABASE

	pgDatabase := db.NewPGDatabase(host, port, user, password, database)
	connection, err = pgDatabase.GetPGPool()

	if err != nil {
		panic("error connecting to database")
	}
}

func GetPostgresConnection() *pgxpool.Pool {
	return connection
}
