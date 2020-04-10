package config_test

import (
	"os"
	"testing"

	"github.com/pocockn/database-migrator/config"
	"github.com/stretchr/testify/assert"
)

func TestConfigCreation(t *testing.T) {
	err := os.Setenv("ENV", "development")
	assert.Nil(t, err)

	expectedConfigStruct := config.Config{
		Database: config.Database{
			Host:         "127.0.0.1",
			DatabaseName: "docker_pocockn",
			Port:         "3306",
			Password:     "pocockn",
			Username:     "pocockn",
		},
	}
	config := config.NewConfig()

	assert.Equal(t, expectedConfigStruct.Database.Host, config.Database.Host)
	assert.Equal(t, expectedConfigStruct.Database.DatabaseName, config.Database.DatabaseName)
	assert.Equal(t, expectedConfigStruct.Database.Port, config.Database.Port)
	assert.Equal(t, expectedConfigStruct.Database.Password, config.Database.Password)
	assert.Equal(t, expectedConfigStruct.Database.Username, config.Database.Username)
}
