package models

// testCon defines the database connection used to test database transactions
var testCon *Connection

// getConnection create a postgres database connection
func getConnection() {
	testCon = InitDB("amity", "amity", "12345")
}

// createSchemas delete old schemam before creating new ones
func createSchemas() error {
	var err error

	if testCon == nil {
		getConnection()
	}

	for _, model := range []interface{}{&UserSpaces{}, &Room{}} {
		testCon.DropTable(model, nil)

		if err = testCon.CreateTable(model, nil); err != nil {
			return err
		}
	}

	return nil
}
