package constants

import (
	"os"

	"github.com/caarlos0/env"
)

// Env is environment variables
var Env envVars

func init() {
	err := env.Parse(&Env)
	if err != nil {
		os.Exit(1)
	}
}

type envVars struct {
	Port string `env:"PORT" envDefault:"3050"`

	// Redis config
	RedisURL string `env:"REDIS_URL" envDefault:"http://localhost:6379"`
}
