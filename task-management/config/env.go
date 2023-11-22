package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	PORT         string
	DATABASE_URL string
}

func (e *Env) Load() error {
	godotenv.Load()
	portErr := loadEnv(&e.PORT, "PORT")
	if portErr != nil {
		return portErr
	}
	databaseUrlErr := loadEnv(&e.DATABASE_URL, "DATABASE_URL")
	if databaseUrlErr != nil {
		return databaseUrlErr
	}
	return nil
}

func loadEnv(value *string, requiredValue string) error {
	*value = os.Getenv(requiredValue)
	if *value == "" {
		return fmt.Errorf("%v env is required", requiredValue)
	}
	return nil
}
