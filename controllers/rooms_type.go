package controllers

import "Amity-Golang/models"

type (
	// Datastore defines all the methods that are implemented by this interface
	Datastore interface {
		CreateRoom(name, roomType string) (models.Room, error)
		CreateUser(fname, lname, officeID, jobType string, livingspaceID ...string) (models.UserSpaces, error)
		DeleteRoom(ID string) (string, error)
		DeleteUser(ID string) (string, error)
		GetRoom(ID string) (models.Room, error)
		GetRooms() ([]models.Room, error)
		GetUser(ID string) (models.UserSpaces, error)
		GetUsers(officeID, livingSpaceID string) ([]models.User, error)
		UpdateRoom(name, ID string) (string, error)
		UpdateUser(fname, lname, ID string) (string, error)
	}

	// Room defines the Room as used by the View part.
	// A Room should contain its room specific details and details of its
	// occupants
	Room struct {
		models.Room
		Occupants []models.User
	}
)
