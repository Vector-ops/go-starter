package configs

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv loads the .env file.
// If .env.production is not found it loads the local development file.
func LoadEnv() {
	dirRetryList := []string{``, `../`, `../../`, `../../../`}
	var err error
	for _, dirPrefix := range dirRetryList {
		envFile := dirPrefix + `.env`
		err = godotenv.Overload(envFile + `.development.local`)
		if err == nil {
			slog.Info(`file .env.development.local loaded (development environment)`)
			return
		}

		err = godotenv.Overload(envFile)
		if err == nil {
			slog.Info(`file .env loaded (production environment)`)
			return
		}
	}
	panic("cannot load .env file")
}

// GetEnv takes two parameters,
// key of the environment variable and
// fallback if the environment variable is not found
func GetEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
