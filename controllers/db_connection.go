package controllers

import (
	"amity-golang/models"
	"fmt"
	"os"

	"github.com/subosito/gotenv"
)

// DBConfig defines the database configuration that is used to establish a database connection
// for the models
var DBConfig *models.PgConfig

// GetConnection create a database connection
func GetConnection(config *models.PgConfig) *models.Connection {
	return models.InitDB(*config)
}

// init sets up the database configuration from the set environment variables
func init() {
	var (
		err error

		envFile = os.Getenv("ENV_FILE")

		exit = func(msg string, err error) {
			msg = msg + " environment variable cannot be empty"

			if err != nil {
				msg = err.Error()
			}

			fmt.Println(msg)
			os.Exit(1)
		}
	)

	DBConfig = &models.PgConfig{}

	if envFile == "" {
		err = gotenv.Load()
	} else {
		err = gotenv.Load(envFile)
	}

	if err != nil {
		exit("", err)
	}

	DBConfig.DBName = os.Getenv("AG_DATABASE")
	DBConfig.DBPassword = os.Getenv("AG_PASSWORD")
	DBConfig.DBUser = os.Getenv("AG_USER")

	switch {
	case DBConfig.DBName == "":
		exit("AG_DATABASE", nil)

	case DBConfig.DBPassword == "":
		exit("AG_PASSWORD", nil)

	case DBConfig.DBUser == "":
		exit("AG_PASSWORD", nil)
	}
}
