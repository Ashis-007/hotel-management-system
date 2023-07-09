package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type Env struct {
	Port       string `mapstructure:"PORT"`
	DBName     string `mapstructure:"DB_NAME"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	JWTSecret  string `mapstructure:"JWT_SECRET"`
	Salt       string `mapstructure:"SALT"`
}

func LoadEnv(path string) (Env, error) {
	envFile := ".env.development"

	envType := os.Getenv("GO_ENV")
	if envType == "production" {
		envFile = ".env.production"
	}

	viper.SetConfigFile(path + envFile)

	var env Env
	if err := viper.Unmarshal(&env); err != nil {
		log.Fatal("failed to load env config: ", err)
		return env, err
	}
	return env, nil
}
