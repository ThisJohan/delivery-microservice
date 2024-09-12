package env

import (
	envParser "github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

func LoadConfig[T any](serviceName string, configStruct *T) error {
	if err := godotenv.Load("./configs/" + serviceName + ".env"); err != nil {
		return err
	}
	if err := envParser.Parse(configStruct); err != nil {
		return err
	}
	return nil
}
