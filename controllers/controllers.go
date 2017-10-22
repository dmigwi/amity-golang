package controllers

import (
	"amity-golang/models"
	"errors"
	"strings"
)

// AddPerson adds a new user to the Amity room allocation system. The new user can
// choose to have a room while bieng created or a room can be allocated to them later
func AddPerson(ds Datastore, fname, lname, jobType, officeName, livingspaceName string) (User, error) {
	var (
		newUser   User
		user, err = ds.GetUser(fname, lname, "")
	)

	if err != nil {
		return newUser, err
	}

	if user != (models.UserSpaces{}) {
		return newUser, errors.New(fname + " " + lname + " already exists")
	}

	if livingspaceName != "" {
		if newUser.LivingSpace, _, err = getRoomDetails(ds, livingspaceName, "livingspace"); err != nil {
			return newUser, err
		}
	}

	if officeName != "" {
		if newUser.Office, _, err = getRoomDetails(ds, officeName, "office"); err != nil {
			return newUser, err
		}
	}

	user, err = ds.CreateUser(fname, lname, jobType, newUser.Office.ID, newUser.LivingSpace.ID)

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
		fetchedRoom.Occupants, err = ds.GetUsers(room.ID, "")
	} else {
		fetchedRoom.Occupants, err = ds.GetUsers("", room.ID)
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
func ReallocatePerson(ds Datastore, fname, lname, roomName string) (Room, error) {
	var (
		status    string
		users     []models.User
		user      models.UserSpaces
		room, err = ds.GetRoom(roomName, "")
	)

	if err != nil {
		return Room{}, err
	}

	if room == (models.Room{}) {
		return Room{}, errors.New(roomName + "does not exist")
	}

	user, err = ds.GetUser(fname, lname, "")

	if err != nil {
		return Room{}, err
	}

	if user == (models.UserSpaces{}) {
		return Room{}, errors.New(fname + " " + lname + " does not exist in the system")
	}

	switch room.Type {
	case "office":
		status, err = ds.UpdateUser(fname, lname, user.ID, room.ID, "")
	case "livingspace":
		status, err = ds.UpdateUser(fname, lname, user.ID, "", room.ID)
	}

	if err != nil {
		return Room{}, err
	}

	if status != "success" {
		return Room{}, errors.New("Room reallocation failed")
	}

	room, users, err = getRoomDetails(ds, room.Name, room.Type)

	return Room{Room: room, Occupants: users}, err
}

// getRoomDetails checks if the room name provided exists, otherwise an error is returned.
// It also check if the existing has reached its capacity, an error is returned if the
// capacity has been reached.
func getRoomDetails(ds Datastore, roomName, roomType string) (models.Room, []models.User, error) {
	var (
		users     []models.User
		room, err = ds.GetRoom(roomName, "")
	)

	if err != nil {
		return room, users, err
	}

	if room == (models.Room{}) {
		return room, users, errors.New(roomName + " " + roomType + " does not exists")
	}

	switch strings.ToLower(roomType) {
	case "office":
		users, err = ds.GetUsers(room.ID, "")
	case "livingspace":
		users, err = ds.GetUsers("", room.ID)
	}

	if err != nil {
		return room, users, err
	}

	if len(users) >= room.Capacity {
		return room, users, errors.New(roomName + " has already reached its maximum capacity")
	}

	return room, users, nil
}
