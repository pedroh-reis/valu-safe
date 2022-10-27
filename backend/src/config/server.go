package config

import (
	"fmt"
	"os"
)

type ServerConfig struct {
	HOST string
	PORT string
}

func NewServerConfig() ServerConfig {
	return ServerConfig{
		HOST: os.Getenv("HOST"),
		PORT: os.Getenv("PORT"),
	}
}

func (I ServerConfig) GetAddress() string {
	return fmt.Sprintf("%s:%s", I.HOST, I.PORT)
}
