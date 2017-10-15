package models

import (
	"errors"
	"strings"

	"github.com/google/uuid"
)

// Room defines the underlying information of a room in Amity room allocation system
type Room struct {
	Capacity int
	ID       string
	Name     string
	Type     string
}

// CreateRoom makes a database operation to create a new room of the provided type
func (config *Connection) CreateRoom(name, roomType string) (Room, error) {
	var capacity int

	switch strings.ToLower(roomType) {
	case "office":
		capacity = 4
	case "livingspace":
		capacity = 6
	default:
		return Room{}, errors.New("Only office and livingspace room types can be created")
	}

	var room = Room{
		Capacity: capacity,
		Type:     roomType,
		Name:     name,
		ID:       uuid.New().String(),
	}

	return room, config.Insert(&room)
}

// DeleteRoom deletes a room given the its ID
func (config *Connection) DeleteRoom(ID string) (string, error) {
	var resp, err = config.Model(&Room{ID: ID}).Delete()

	if err != nil {
		return "error", err
	}

	if resp.RowsAffected() == 1 {
		return "success", nil
	}

	return "error", errors.New("No Room that was deleted")

}

// GetRoom fetches the room details given its ID
func (config *Connection) GetRoom(ID string) (Room, error) {
	var room = Room{ID: ID}

	return room, config.Select(&room)
}

// GetRooms fetches all rooms that are currently in existence
func (config *Connection) GetRooms() ([]Room, error) {
	var rooms []Room

	return rooms, config.Model(&rooms).Select()

}

// UpdateRoom updates the name of a given room
func (config *Connection) UpdateRoom(name, ID string) (string, error) {
	var resp, err = config.Model(&Room{Name: name, ID: ID}).Column("name").Update()

	if err != nil {
		return "error", err
	}

	if resp.RowsAffected() == 1 {
		return "success", nil
	}

	return "error", errors.New("No Room that was updated")
}
