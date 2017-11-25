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
		LivingSpaceID string
		OfficeID      string
	}
)

// CreateUser create a new user with the provided details
func (config *Connection) CreateUser(fname, lname, userType, officeID, livingspaceID string) (UserSpaces, error) {

	switch strings.ToLower(userType) {
	case "fellow", "staff":
		userType = strings.ToLower(userType)

	default:
		return UserSpaces{}, errors.New(userType + ": Only staff and fellow userTypes are allowed")
	}

	if livingspaceID != "" && userType == "staff" {
		return UserSpaces{}, errors.New("A staff cannot be allocated a livingspace")
	}

	var user = UserSpaces{
		User: User{
			FirstName: fname,
			ID:        uuid.New().String(),
			LastName:  lname,
			Type:      userType,
		},
		LivingSpaceID: livingspaceID,
		OfficeID:      officeID,
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
	var (
		err  error
		user UserSpaces
	)

	switch {
	case fname != "" && lname != "":
		err = config.Model(&user).Where("first_name =?", fname).Where("last_name =?", lname).Select()
	case ID != "":
		err = config.Model(&user).Where("id =?", ID).Select()
	default:
		err = errors.New("ID or Firstname and Lastname of a user must be provided")
	}

	if err != nil && err.Error() != "pg: no rows in result set" {
		return user, err
	}

	return user, nil
}

// GetUsers fetches all the users currently in existence
func (config *Connection) GetUsers(officeID, livingSpaceID string) ([]User, error) {
	var (
		err      error
		newUsers []User
		users    []UserSpaces
	)

	switch {
	case officeID == "" && livingSpaceID != "":
		err = config.Model(&users).Where("Living_Space_ID = ?", livingSpaceID).Select()

	case livingSpaceID == "" && officeID != "":
		err = config.Model(&users).Where("Office_ID = ?", officeID).Select()

	default:
		err = config.Model(&users).Where("Office_ID = ?", officeID).Where("Living_Space_ID = ?", livingSpaceID).Select()
	}

	if err != nil && err.Error() != "pg: no rows in result set" {
		return newUsers, err
	}

	for _, user := range users {
		newUsers = append(newUsers, user.User)
	}

	return newUsers, nil
}

// UpdateUser updates the firstName and the last name to the user whose ID is provided
func (config *Connection) UpdateUser(fname, lname, ID, officeID, livingspaceID string) (string, error) {
	var user = UserSpaces{User: User{
		FirstName: fname,
		LastName:  lname,
		ID:        ID,
	}}

	if livingspaceID != "" {
		user.LivingSpaceID = livingspaceID
	}

	if officeID != "" {
		user.OfficeID = officeID
	}

	var resp, err = config.Model(&user).Column("first_name", "last_name", "living_space_id", "office_id").Update()

	if err != nil {
		return "error", err
	}

	if resp.RowsAffected() == 1 {
		return "success", nil
	}

	return "error", errors.New("No User that was updated")
}
