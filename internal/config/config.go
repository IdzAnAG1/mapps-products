package config

import (
	"fmt"
	"os"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type ServerConfig struct {
	Port string `env:"PRODUCT_SERVER_PORT" envDefault:"8082"`
}

type Logger struct {
	Level int `env:"LOG_LEVEL" envDefault:"0"`
}

type DatabaseConfig struct {
	Host     string `env:"PRODUCT_DATABASE_HOST"     envDefault:"localhost"`
	Port     int    `env:"PRODUCT_DATABASE_PORT"     envDefault:"5432"`
	User     string `env:"PRODUCT_DATABASE_USER"     envDefault:"postgres"`
	Password string `env:"PRODUCT_DATABASE_PASSWORD" envDefault:"postgres"`
	Name     string `env:"PRODUCT_DATABASE_NAME"     envDefault:"mapps_products_db"`
}

func (d *DatabaseConfig) DSN() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		d.User, d.Password, d.Host, d.Port, d.Name,
	)
}

type Config struct {
	Server   ServerConfig
	Logger   Logger
	Database DatabaseConfig
}

func LoadAndGetConfig() (*Config, error) {
	err := godotenv.Load(os.Getenv("ENV_FILE"))
	if err != nil {
		fmt.Printf(
			"Error loading env file: %v \n"+
				"Trying to read environment variables\n", err,
		)
	}

	config := &Config{}

	err = env.Parse(config)
	if err != nil {
		fmt.Println("Failed to parse env variables")
		return nil, err
	}

	return config, nil
}
