package models

import (
	"errors"
	"strings"
)

// Room defines the underlying information of a room in Amity room allocation system
type Room struct {
	Capacity int
	ID       string
	Name     string
	Type     string
}

// CreateRoom makes a database operation to create a new room of the provided type
func (config *Connection) CreateRoom(name, roomType string, size ...int) (Room, error) {
	var (
		capacity int

		room = Room{
			Capacity: capacity,
			Type:     roomType,
			Name:     name,
		}
	)

	switch strings.ToLower(roomType) {
	case "office":
		capacity = 4
	case "livingspace":
		capacity = 6
	default:
		return room, errors.New("Only office and livingspace room types can be created")
	}

	return room, config.Insert(&room)
}

// DeleteRoom deletes a room given the its ID
func (config *Connection) DeleteRoom(ID string) error {
	return config.Delete(&Room{
		ID: ID,
	})
}

// GetRoom fetches the room details given its ID
func (config *Connection) GetRoom(ID string) (Room, error) {
	var room Room

	return room, config.Select(&room)
}

// GetRooms fetches all rooms that are currently in existence
func (config *Connection) GetRooms() ([]Room, error) {
	var rooms []Room

	return rooms, config.Model(&rooms).Select()

}

// UpdateRoom updates the name of a given room
func (config *Connection) UpdateRoom(name string) (Room, error) {
	var room = Room{Name: name}

	return room, config.Update(&room)
}
