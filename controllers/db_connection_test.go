package controllers

import (
	"log"
	"os"
	"testing"
)

var con Datastore

// TestMain sets up the test environment
func TestMain(m *testing.M) {
	if !testing.Verbose() {
		log.SetFlags(0)
		// log.SetOutput(ioutil.Discard)
	}

	switch os.Getenv("ENV") {
	case "DB":
		con = GetConnection(DBConfig)

	case "mockDB":
		con = &MockConn{}
	}

	os.Exit(m.Run())
}
