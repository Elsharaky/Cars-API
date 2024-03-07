package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

func LoadENV() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	if os.Getenv("API_HOST") == "" {
		return errors.New("API_HOST not set in .env file")
	}

	if os.Getenv("API_PORT") == "" {
		return errors.New("API_PORT not set in.env file")
	}

	if os.Getenv("DB_HOST") == "" {
		return errors.New("DB_HOST not set in.env file")
	}

	if os.Getenv("DB_PORT") == "" {
		return errors.New("DB_PORT not set in.env file")
	}

	if os.Getenv("DB_NAME") == "" {
		return errors.New("DB_NAME not set in.env file")
	}

	if os.Getenv("DB_USER") == "" {
		return errors.New("DB_USER not set in.env file")
	}

	if os.Getenv("DB_PASSWORD") == "" {
		return errors.New("DB_PASSWORD not set in.env file")
	}

	return nil
}
