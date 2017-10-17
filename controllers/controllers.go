package controllers

import (
	"Amity-Golang/models"
	"errors"
)

// AddPerson adds a new user to the Amity room allocation system. The new user can
// choose to have a room while bieng created or a room can be allocated to them later
func AddPerson(ds Datastore, fname, lname, jobType, office, livingspace string) (User, error) {
	var (
		newUser   User
		user, err = ds.GetUser(fname, lname, "")
	)

	if err != nil {
		return newUser, err
	}

	if user != (models.UserSpaces{}) {
		newUser.Office, err = ds.GetRoom("", user.Office)

		if err != nil {
			return newUser, err
		}

		newUser.LivingSpace, err = ds.GetRoom("", user.Livingspace)

		if err != nil {
			return newUser, err
		}

		newUser.User = user.User

		return newUser, nil
	}

	if livingspace != "" {
		if newUser.LivingSpace, err = ds.GetRoom(livingspace, ""); err != nil {
			return newUser, err
		}

		if newUser.LivingSpace == (models.Room{}) {
			return newUser, errors.New("LivingSpace :" + livingspace + " does not exists")
		}
	}

	if office != "" {
		if newUser.Office, err = ds.GetRoom(office, ""); err != nil {
			return newUser, err
		}

		if newUser.Office == (models.Room{}) {
			return newUser, errors.New("Office :" + office + " does not exists")
		}

	}

	user, err = ds.CreateUser(fname, lname, newUser.Office.ID, jobType, newUser.LivingSpace.ID)

	newUser.User = user.User

	return newUser, err

}

// CreateRoom creates a room of the specified type if no other room with the same name exists
func CreateRoom(ds Datastore, name, roomType string) (Room, error) {
	var room, err = ds.GetRoom(name, "")

	if err != nil {
		return Room{}, err
	}

	if room != (models.Room{}) {
		return Room{Room: room}, nil
	}

	room, err = ds.CreateRoom(name, roomType)

	return Room{Room: room}, err
}

// GetRoom retrieves a room and all its occupants if it exists
func GetRoom(ds Datastore, name string) (Room, error) {
	var (
		fetchedRoom Room
		room, err   = ds.GetRoom(name, "")
	)

	if err != nil {
		return fetchedRoom, err
	}

	if room == (models.Room{}) {
		return fetchedRoom, errors.New("Room :" + name + " does not exist")
	}

	fetchedRoom.Room = room

	if room.Type == "office" {
		fetchedRoom.Occupants, err = ds.GetUsers(name, "")
	} else {
		fetchedRoom.Occupants, err = ds.GetUsers("", name)
	}

	return fetchedRoom, err
}

// GetRoomAllocations fetches all rooms with their current occupants
func GetRoomAllocations(ds Datastore) ([]Room, error) {
	var (
		room Room

		fetchedRooms = make([]Room, 0)
		rooms, err   = ds.GetRooms()
	)

	if err != nil {
		return fetchedRooms, err
	}

	for _, rm := range rooms {
		room, err = GetRoom(ds, rm.Name)

		if err != nil {
			return fetchedRooms, err
		}

		fetchedRooms = append(fetchedRooms, room)
	}

	return fetchedRooms, err
}

// GetUnallocatedPeople fetches and returns all users don't have any room allocated to them
func GetUnallocatedPeople(ds Datastore) ([]User, error) {
	var (
		users, err = ds.GetUsers("", "")
		usrs       = make([]User, 0)
	)

	if err != nil {
		return usrs, err
	}

	for _, usr := range users {
		usrs = append(usrs, User{User: usr})
	}

	return usrs, err
}

// ReallocatePerson assign a user already added to the system a new room. It can also move a
// a user to a new room.
func ReallocatePerson(Datastore interface{}, userID, roomName string) (Room, error) {

}
