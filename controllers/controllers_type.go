package controllers

import (
	"amity-golang/models"
)

type (
	// Datastore defines all the methods that are implemented by this interface
	Datastore interface {
		CreateRoom(name, roomType string) (models.Room, error)
		CreateUser(fname, lname, jobType, officeID, livingspaceID string) (models.UserSpaces, error)
		DeleteRoom(ID string) (string, error)
		DeleteUser(ID string) (string, error)
		DestroyData() error
		GetRoom(name, ID string) (models.Room, error)
		GetRooms() ([]models.Room, error)
		GetUser(fname, lname, ID string) (models.UserSpaces, error)
		GetUsers(officeID, livingSpaceID string) ([]models.User, error)
		UpdateRoom(name, ID string) (string, error)
		UpdateUser(fname, lname, ID, officeID, livingspaceID string) (string, error)
	}

	// Room defines the Room as used by the View part.
	// A Room should contain its room specific details and details of its
	// occupants
	Room struct {
		models.Room
		Occupants []models.User
	}

	// User defines a user of the amity room allocation system as defined in the models
	User struct {
		models.User
		Office      models.Room
		LivingSpace models.Room
	}
)
