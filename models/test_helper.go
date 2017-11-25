package models

import (
	"errors"
	"os"
)

var (
	// testCon defines the database connection used to test database transactions
	testCon *Connection

	// pgConfig defines a PgConfig instance used to hold postgres test db configuration
	pgConfig *PgConfig
)

// getDBConfig create a postgres database configuration using the set environment variables
func getDBConfig() (PgConfig, error) {
	var dbConfig PgConfig

	dbConfig.DBName = os.Getenv("AG_DATABASE")
	dbConfig.DBPassword = os.Getenv("AG_PASSWORD")
	dbConfig.DBUser = os.Getenv("AG_USER")

	if dbConfig.DBName == "" || dbConfig.DBPassword == "" || dbConfig.DBUser == "" {
		return PgConfig{}, errors.New("AG_DATABASE, AG_PASSWORD or AG_USER " +
			"environment variable cannot be empty")
	}

	return dbConfig, nil
}

// CreateSchemas delete old schemam before creating new ones
func CreateSchemas() error {
	var config, err = getDBConfig()

	if err != nil {
		return err
	}

	testCon = InitDB(config)

	for _, model := range []interface{}{&UserSpaces{}, &Room{}} {
		testCon.DropTable(model, nil)

		if err = testCon.CreateTable(model, nil); err != nil {
			return err
		}
	}

	return nil
}
