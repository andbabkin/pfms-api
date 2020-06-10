package config

import (
	"os"

	"github.com/joho/godotenv"
)

// LoadEnvIntoOS loads environement variables from files into OS
func LoadEnvIntoOS(filenames ...string) error {
	envMap, err := godotenv.Read(filenames...)
	if err != nil {
		return err
	}

	for key, value := range envMap {
		err = os.Setenv(key, value)
		if err != nil {
			return err
		}
	}

	return nil
}
