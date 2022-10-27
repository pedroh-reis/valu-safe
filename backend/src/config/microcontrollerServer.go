package config

import (
	"fmt"
	"os"
)

type MicrocontrollerServerConfig struct {
	HOST string
	PORT string
}

func NewMicrocontrollerServerConfig() MicrocontrollerServerConfig {
	return MicrocontrollerServerConfig{
		HOST: os.Getenv("MICROCONTROLLER_HOST"),
		PORT: os.Getenv("MICROCONTROLLER_PORT"),
	}
}

func (I MicrocontrollerServerConfig) GetAddress() string {
	return fmt.Sprintf("%s:%s", I.HOST, I.PORT)
}
