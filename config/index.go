package config

import (
	"fmt"
	"os"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
    Server struct {
        Port string `envconfig:"SERVER_PORT"`
        Host string `envconfig:"SERVER_HOST"`
    }
}

var cfg Config

func ReadEnv() {
    err := envconfig.Process("", &cfg)
    if err != nil {
        processError(err)
    }
}


func processError(err error) {
	fmt.Println(`Error processing config: ` + err.Error())
	os.Exit(2)
}

func GetConfig() Config {
	ReadEnv()
	return cfg
}
