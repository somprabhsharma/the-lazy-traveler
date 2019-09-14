package constants

import (
	"os"

	"github.com/caarlos0/env"
)

var Env envVars

func init() {
	err := env.Parse(&Env)
	if err != nil {
		os.Exit(1)
	}
}

type envVars struct {
	Port              string `env:"PORT" envDefault:"3050"`
}
