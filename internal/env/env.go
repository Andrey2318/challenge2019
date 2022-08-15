package env

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"sync"
)

var onceLoadEnv sync.Once

func LoadEnvs() {
	onceLoadEnv.Do(func() {
		envFilename := ".env"
		if v := os.Getenv("ENV_FILE"); v != "" {
			envFilename = v
		}
		if err := godotenv.Load(envFilename); err != nil {
			log.Panic(err)
		}
	})
}

func MustEnv(key string) string {
	if os.Getenv(key) == "" {
		log.Panicf("unknown %s param in ENV", key)
	}
	return os.Getenv(key)
}

func Env(key string, defaultValue string) string {
	if os.Getenv(key) == "" {
		return defaultValue
	}
	return os.Getenv(key)
}
