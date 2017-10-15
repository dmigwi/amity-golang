package models

import (
	"log"
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// TestMain runs before all tests run
func TestMain(m *testing.M) {
	var err = createSchemas()

	if !testing.Verbose() {
		log.SetFlags(0)
	}

	if err != nil {
		panic(err)
	}

	os.Exit(m.Run())
}

// TestInitDB tests the functionality of InitDB
func TestInitDB(t *testing.T) {
	var con *Connection
	Convey("Tests for InitDB ", t, func() {
		Convey("The connection created should initially be nil", func() {
			So(con, ShouldBeNil)
		})

		Convey("The connection should not be nil after being initialzed", func() {
			con = InitDB("amity", "amity", "12345")

			So(con, ShouldNotBeNil)
		})

	})
}
