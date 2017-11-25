package models

import "github.com/go-pg/pg"

type (
	// Connection defines the database connection that is needed to make a
	// database transaction
	Connection struct {
		*pg.DB
	}

	// PgConfig defines the configuration needed to connect to a postgres database instance
	PgConfig struct {
		DBUser     string
		DBName     string
		DBPassword string
	}
)

// InitDB creates a database connection that is used to execute
// the various database transactions
func InitDB(config PgConfig) *Connection {
	var con = pg.Connect(
		&pg.Options{
			User:     config.DBUser,
			Database: config.DBName,
			Password: config.DBPassword,
		})

	return &Connection{con}
}

// DestroyData recreates the database thus destroying
// all the data that existed previously
func (config *Connection) DestroyData() error {
	return CreateSchemas()
}
