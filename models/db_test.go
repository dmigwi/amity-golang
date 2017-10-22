package models

import (
	"io/ioutil"
	"log"
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// TestMain sets up the test environment
func TestMain(m *testing.M) {
	var err = CreateSchemas()

	if err != nil {
		log.Println("Failed to initialize the DB :", err.Error())
		os.Exit(1)
	}

	if !testing.Verbose() {
		log.SetFlags(0)
		log.SetOutput(ioutil.Discard)
	}

	os.Exit(m.Run())
}

// TestInitDB tests the functionality of InitDB
func TestInitDB(t *testing.T) {
	var (
		con        *Connection
		configType PgConfig

		getConn = func() {
			var config, err = getDBConfig()

			So(err, ShouldBeNil)
			So(config, ShouldHaveSameTypeAs, configType)
			So(config, ShouldNotResemble, configType)

			con = InitDB(config)
		}
	)

	Convey("Tests for InitDB ", t, func() {
		Convey("The connection created should initially be nil", func() {
			So(con, ShouldBeNil)
		})

		Convey("The connection should not be nil after being initialized", func() {
			getConn()

			So(con, ShouldNotBeNil)
		})

	})
}
