package models

import (
	"errors"
	"strings"

	"github.com/google/uuid"
)

type (
	// User defines basic details about a user. They can be used to identify a user
	User struct {
		FirstName string
		ID        string
		LastName  string
		Type      string
	}

	// UserSpaces defines details of the occupants of the room in amity space allocation system
	UserSpaces struct {
		User
		Livingspace string
		Office      string
	}
)

// CreateUser create a new user with the provided details
func (config *Connection) CreateUser(fname, lname, office, jobType string, livingspace ...string) (UserSpaces, error) {
	var ls = ""

	if len(livingspace) > 0 {
		ls = livingspace[0]
	}

	switch strings.ToLower(jobType) {
	case "fellow", "staff":
		jobType = strings.ToLower(jobType)

	default:
		return UserSpaces{}, errors.New(jobType + ": Only staff and fellow jobTypes are allowed")
	}

	if ls != "" && jobType == "staff" {
		return UserSpaces{}, errors.New("A staff cannot be allocated a livingspace")
	}

	var user = UserSpaces{
		User: User{
			FirstName: fname,
			ID:        uuid.New().String(),
			LastName:  lname,
			Type:      jobType,
		},
		Livingspace: ls,
		Office:      office,
	}

	return user, config.Insert(&user)
}

// DeleteUser deletes the user given their user  ID
func (config *Connection) DeleteUser(ID string) (string, error) {
	var resp, err = config.Model(&UserSpaces{User: User{ID: ID}}).Delete()

	if err != nil {
		return "error", err
	}

	if resp.RowsAffected() == 1 {
		return "success", nil
	}

	return "error", errors.New("No user that was deleted")
}

// GetUser fetches and returns a user associated with the given ID
func (config *Connection) GetUser(fname, lname, ID string) (UserSpaces, error) {
	var user UserSpaces

	switch {
	case fname != "" && lname != "":
		return user, config.Model(&user).Where("first_name =?", fname).Where("last_name =?", lname).Select()
	case ID != "":
		return user, config.Select(&UserSpaces{User: User{ID: ID}})
	default:
		return user, errors.New("ID or Firstname and Lastname of a user must be provided")
	}
}

// GetUsers fetches all the users currently in existence
func (config *Connection) GetUsers(office, livingSpace string) ([]User, error) {
	var (
		err      error
		newUsers []User
		users    []UserSpaces
	)

	switch {
	case office == "" && livingSpace != "":
		err = config.Model(&users).Where("Livingspace = ?", livingSpace).Select()

	case livingSpace == "" && office != "":
		err = config.Model(&users).Where("Office = ?", office).Select()

	default:
		err = config.Model(&users).Where("Office = ?", office).Where("Livingspace = ?", livingSpace).Select()
	}

	if err != nil {
		return newUsers, err
	}

	for _, user := range users {
		newUsers = append(newUsers, user.User)
	}

	return newUsers, nil
}

// UpdateUser updates the firstName and the last name to the user whose ID is provided
func (config *Connection) UpdateUser(fname, lname, ID string) (string, error) {
	var resp, err = config.Model(&UserSpaces{User: User{
		FirstName: fname,
		LastName:  lname,
		ID:        ID,
	}}).Update()

	if err != nil {
		return "error", err
	}

	if resp.RowsAffected() == 1 {
		return "success", nil
	}

	return "error", errors.New("No User that was updated")
}
