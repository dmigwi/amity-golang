package models

import "github.com/go-pg/pg"

// Connection defines the database connection that is needed to make a
// database transaction
type Connection struct {
	*pg.DB
}

// InitDB creates a database connection that is used to execute
// the various database transactions
func InitDB(user, db, password string) *Connection {
	var con = pg.Connect(
		&pg.Options{
			User:     user,
			Database: db,
			Password: password,
		})

	return &Connection{con}
}
