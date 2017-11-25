package controllers

import (
	"amity-golang/models"
	"errors"
	"log"

	"github.com/google/uuid"
)

// MockConn defines a struct that helps in mocking the models package
type MockConn struct {
	Name string
}

var (
	// detailedUsers defines a slice that holds user details thereby mocking the UserSpaces model
	detailedUsers = make([]*models.UserSpaces, 0)

	// rooms defines a slice that holds room details thereby mocking the Room model
	rooms = make([]*models.Room, 0)
)

// CreateRoom mocks the CreateRoom method in the models subpackage
func (config *MockConn) CreateRoom(name, roomType string) (models.Room, error) {
	var capacity = 6

	if roomType == "livingspace" {
		capacity = 4
	}

	return generateRoom(capacity, name, roomType), nil
}

// CreateUser mocks the CreateUser method in the models subpackage
func (config *MockConn) CreateUser(fname, lname, jobType, officeID, livingspaceID string) (models.UserSpaces, error) {
	return generateDetailedUser(fname, lname, jobType, officeID, livingspaceID), nil
}

// DeleteRoom mocks the DeleteRoom method in the models subpackage
func (config *MockConn) DeleteRoom(ID string) (string, error) {
	return "success", nil
}

// DestroyData destroys all the test data written into detailedUser and rooms slices
func (config *MockConn) DestroyData() error {
	detailedUsers = make([]*models.UserSpaces, 0)

	rooms = make([]*models.Room, 0)

	return nil
}

// DeleteUser mocks the DeleteUser method in the models subpackage
func (config *MockConn) DeleteUser(ID string) (string, error) {
	return "success", nil
}

// GetRoom mocks the GetRoom method in the models subpackage
func (config *MockConn) GetRoom(name, ID string) (models.Room, error) {
	if name != "" {
		for _, val := range rooms {
			if name == val.Name {
				return *val, nil
			}
		}
	} else if ID != "" {
		for _, val := range rooms {
			if name == val.ID {
				return *val, nil
			}
		}
	}
	return models.Room{}, nil
}

// GetRooms mocks the functionality of GetRooms method in the models subpackage
func (config *MockConn) GetRooms() ([]models.Room, error) {
	var rms []models.Room

	for _, val := range rooms {
		rms = append(rms, *val)
	}

	return rms, nil
}

// GetUser mocks the functionality of GetUser method in the models subpackage
func (config *MockConn) GetUser(fname, lname, ID string) (models.UserSpaces, error) {
	if fname != "" && lname != "" {
		for _, val := range detailedUsers {
			if fname == val.FirstName && lname == val.LastName {
				return *val, nil
			}
		}
	} else if ID != "" {
		for _, val := range detailedUsers {
			if ID == val.ID {
				return *val, nil
			}
		}
	}

	return models.UserSpaces{}, nil
}

// GetUsers mocks the functionality of GetUsers method in the models subpackage
func (config *MockConn) GetUsers(officeID, livingSpaceID string) ([]models.User, error) {
	var usrs []models.User

	if officeID != "" && livingSpaceID != "" {
		for _, val := range detailedUsers {
			if officeID == val.OfficeID && livingSpaceID == val.LivingSpaceID {
				usrs = append(usrs, val.User)
			}
		}
	} else if officeID == "" && livingSpaceID != "" {
		for _, val := range detailedUsers {
			if livingSpaceID == val.LivingSpaceID {
				usrs = append(usrs, val.User)
			}
		}
	} else if officeID != "" && livingSpaceID == "" {
		for _, val := range detailedUsers {
			if officeID == val.OfficeID {
				usrs = append(usrs, val.User)
			}
		}
	}
	return usrs, nil
}

// UpdateRoom mocks the functionality of UpdateRoom method in the models subpackage
func (config *MockConn) UpdateRoom(name, ID string) (string, error) {
	var val *models.Room

	for _, val = range rooms {
		if val.ID == ID {
			break
		} else {
			return "error", errors.New("Room not found")
		}
	}

	val.Name = name

	return "success", nil
}

// UpdateUser mocks the functionality of UpdateUser method in the models subpackage
func (config *MockConn) UpdateUser(fname, lname, ID, officeID, livingspaceID string) (string, error) {
	for index, val := range detailedUsers {

		if val.ID == ID {
			val.FirstName = fname
			val.LastName = lname
			val.ID = ID

			if officeID != "" {
				val.OfficeID = officeID
			}

			if livingspaceID != "" {
				val.LivingSpaceID = livingspaceID
			}

			detailedUsers[index:][0] = val
		}

		return "success", nil
	}

	return "error", errors.New("Room not found")
}

// generateRoom returns a newly created room
func generateRoom(capacity int, name, roomType string) models.Room {
	var room = models.Room{
		Capacity: capacity,
		ID:       uuid.New().String(),
		Name:     name,
		Type:     roomType,
	}
	rooms = append(rooms, &room)

	return room
}

// generateUser returns a newly created user
func generateUser(fname, lname, jobType string) models.User {
	return models.User{
		FirstName: fname,
		ID:        uuid.New().String(),
		LastName:  lname,
		Type:      jobType,
	}
}

// generateDetailedUser returns a newly created user and the office and livingspace allocated to them
func generateDetailedUser(fname, lname, jobtype, officeID, livingspaceID string) models.UserSpaces {
	var user = models.UserSpaces{
		LivingSpaceID: livingspaceID,
		OfficeID:      officeID,
		User:          generateUser(fname, lname, jobtype),
	}

	detailedUsers = append(detailedUsers, &user)

	return user
}

// printRooms prints to the console the contents of the room slice
func printRooms() {
	for index, elem := range rooms {
		log.Println()
		log.Println("Room Count :", index+1)
		log.Println("Name :", elem.Name)
		log.Println("ID :", elem.ID)
		log.Println("Type :", elem.Type)
		log.Println("Capacity :", elem.Capacity)
		log.Println()
	}
}

// printUsers prints to the console the contents of the detailedUsers slice
func printUsers() {
	for index, elem := range detailedUsers {
		log.Println()
		log.Println("User Count :", index+1)
		log.Println("FirstName :", elem.FirstName)
		log.Println("LastName :", elem.LastName)
		log.Println("ID :", elem.ID)
		log.Println("Type :", elem.Type)
		log.Println("Office :", elem.OfficeID)
		log.Println("LivingSpace :", elem.LivingSpaceID)
		log.Println()
	}
}
