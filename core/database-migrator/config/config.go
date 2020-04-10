package config

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

type (
	// Config contains config for the application.
	Config struct {
		Database Database
	}

	// Database holds database values in our config.
	Database struct {
		Host         string
		DatabaseName string
		Port         string
		Password     string
		Username     string
		URL          string
	}
)

// NewConfig creates a new config struct.
func NewConfig() Config {
	var config Config
	if _, err := toml.DecodeFile(config.generatePath(), &config); err != nil {
		fmt.Println(err)
	}

	config.Database.URL = generateURL(
		config.Database.Username,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.DatabaseName,
	)

	return config
}

func generateURL(username string, password string, host string, port string, database string) string {
	templateString := "%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4"

	return fmt.Sprintf(
		templateString,
		username,
		password,
		host,
		port,
		database,
	)
}

func (c Config) environment() string {
	environment := "development"

	if os.Getenv("ENV") != "" {
		environment = os.Getenv("ENV")
	}

	return environment
}

func (c Config) generatePath() string {
	if os.Getenv("ENV") == "test" {
		return "development.toml"
	}

	return fmt.Sprintf("config/%s.toml", c.environment())
}
