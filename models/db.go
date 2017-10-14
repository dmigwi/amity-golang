package models

import "github.com/go-pg/pg"

// Connection defines the database connection that is needed to make a
// database transaction
type Connection struct {
	*pg.DB
}

// InitDB creates a database connection that is used to make
// the various database transactions
func InitDB(config string) (*Connection, error) {
	var db = pg.Connect(
		&pg.Options{
			User:     "amity",
			Database: "amity",
			Password: "12345",
		})

	return &Connection{db}, nil
}
