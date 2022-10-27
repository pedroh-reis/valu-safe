package main

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pedroh-reis/valu-safe/backend/src/config"
)

func GetPostgresDbConnection() *sqlx.DB {
	postgresConfig := config.NewPostgresConfig()

	db, err := sqlx.Connect("postgres", postgresConfig.GetDataSourceName())
	if err != nil {
		log.Fatal(err)
	}

	return db
}
