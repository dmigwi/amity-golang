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
		errStr = " environment variable cannot be empty"

		exit = func(msg string) {
			fmt.Println(msg + errStr)
			os.Exit(1)
		}
	)

	gotenv.Load()

	var db = os.Getenv("AG_DATABASE")

	if db == "" {
		exit("AG_DATABASE")
	}

	var password = os.Getenv("AG_PASSWORD")

	if password == "" {
		exit("AG_PASSWORD")
	}

	var user = os.Getenv("AG_USER")

	if user == "" {
		exit("AG_PASSWORD")
	}

	DBConfig = &models.PgConfig{
		DBName:     db,
		DBPassword: password,
		DBUser:     user,
	}

}
