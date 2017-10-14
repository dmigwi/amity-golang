package models

// User defines details of the occupants of the room in amity space allocation system
type User struct {
	FirstName string
	ID        string
	LastName  string
	Type      string
}

// CreateUser create a new user with the provided details
func (config *Connection) CreateUser(fname, lname, jobType string) (User, error) {
	var user = User{
		FirstName: fname,
		LastName:  lname,
		Type:      jobType,
	}

	return user, config.Insert(&user)
}

// DeleteUser deletes the user given their user  ID
func (config *Connection) DeleteUser(ID string) error {
	var user = User{ID: ID}

	return config.Delete(&user)
}

// GetUser fetches and returns a user associated with the given ID
func (config *Connection) GetUser(ID string) (User, error) {
	var user = User{ID: ID}

	return user, config.Select(&user)
}

// GetUsers fetches all the users currently in existence
func (config *Connection) GetUsers() ([]User, error) {
	var users []User

	return users, config.Model(&users).Select()
}

// UpdateUser updates the firstName and the last name to the user whose ID is provided
func (config *Connection) UpdateUser(fname, lname, ID string) (User, error) {
	var user = User{
		FirstName: fname,
		LastName:  lname,
		ID:        ID,
	}

	return user, config.Update(&user)
}
