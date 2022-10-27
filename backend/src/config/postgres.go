package config

import (
	"fmt"
	"os"
)

type PostgresConfig struct {
	HOST     string
	PORT     string
	DB_NAME  string
	USER     string
	PASSWORD string
}

func NewPostgresConfig() PostgresConfig {
	return PostgresConfig{
		HOST:     os.Getenv("POSTGRES_HOST"),
		PORT:     os.Getenv("POSTGRES_PORT"),
		DB_NAME:  os.Getenv("POSTGRES_DB_NAME"),
		USER:     os.Getenv("POSTGRES_USER"),
		PASSWORD: os.Getenv("POSTGRES_PASSWORD"),
	}
}

func (I PostgresConfig) GetDataSourceName() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		I.HOST, I.PORT, I.USER, I.PASSWORD, I.DB_NAME)
}
