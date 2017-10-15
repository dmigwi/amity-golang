package models

import (
	"log"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// printRoom prints to the console the underlying values of a room
var printRoom = func(rm Room) {
	log.Println()
	log.Println("Capacity :", rm.Capacity)
	log.Println("ID :", rm.ID)
	log.Println("Name :", rm.Name)
	log.Println("Type :", rm.Type)
	log.Println()
}

// TestCreateRooms tests the functionality of CreateRoom
func TestCreateRooms(t *testing.T) {
	var (
		err            error
		room, testRoom Room
		status         string

		fetchRoom = func(name, roomType string) {
			room, err = testCon.CreateRoom(name, roomType)

			So(err, ShouldBeNil)

			status, err = testCon.DeleteRoom(room.ID)

			So(status, ShouldEqual, "success")
		}
	)

	Convey("Tests for CreateRoom Method", t, func() {
		Convey("CreateRoom should return Room values that are not empty and a nil error", func() {
			fetchRoom("PHP", "livingspace")

			So(err, ShouldBeNil)
			So(room, ShouldHaveSameTypeAs, testRoom)
			So(room, ShouldNotResemble, testRoom)
			So(room.Capacity, ShouldEqual, 6)
			So(room.ID, ShouldNotBeEmpty)
			So(room.Name, ShouldEqual, "PHP")
			So(room.Type, ShouldEqual, "livingspace")
		})

		Convey("Print the succesfully created Room", func() {
			fetchRoom("Oculus", "office")

			So(err, ShouldBeNil)

			printRoom(room)

		})
	})
}

// TestDeleteRoom tests the functionality of DeleteRoom
func TestDeleteRoom(t *testing.T) {
	var (
		err    error
		room   Room
		status string

		deleteRoom = func(name, roomType string) {
			room, err = testCon.CreateRoom(name, roomType)

			So(err, ShouldBeNil)
			So(room.ID, ShouldNotBeEmpty)

			status, err = testCon.DeleteRoom(room.ID)

		}
	)

	Convey("Tests for DeleteRoom", t, func() {
		Convey("DeleteRoom should return a status success and a nil error", func() {
			deleteRoom("Krypton", "Office")

			So(err, ShouldBeNil)
			So(status, ShouldEqual, "success")
		})

		Convey("Print successfully fetched status ", func() {
			deleteRoom("Krypton", "Office")

			log.Println()
			log.Println("Status :", status)
			log.Println()
		})
	})
}

// TestGetRoom tests the functionality of GetRoom
func TestGetRoom(t *testing.T) {
	var (
		err           error
		room, newRoom Room
		status        string

		fetchRoom = func(name, roomType string) {
			room, err = testCon.CreateRoom(name, roomType)

			So(err, ShouldBeNil)
			So(room.ID, ShouldNotBeEmpty)

			newRoom, err = testCon.GetRoom(room.ID)

			So(err, ShouldBeNil)

			status, err = testCon.DeleteRoom(room.ID)

			So(status, ShouldEqual, "success")
		}
	)

	Convey("Tests for GetRoom", t, func() {
		Convey("GetRoom should return Room values that are not empty", func() {
			fetchRoom("Yorkshire", "livingspace")

			So(err, ShouldBeNil)
			So(newRoom.Capacity, ShouldEqual, 6)
			So(newRoom.ID, ShouldNotBeEmpty)
			So(newRoom.Name, ShouldEqual, "Yorkshire")
			So(newRoom.Type, ShouldEqual, "livingspace")
		})

		Convey("Print the successfully fetched Room", func() {
			fetchRoom("Mogadishu", "livingspace")

			So(err, ShouldBeNil)

			printRoom(newRoom)
		})
	})
}

// TestGetRooms tests the functionality of GetRooms
func TestGetRooms(t *testing.T) {
	var (
		err    error
		room   Room
		rooms  []Room
		status string

		fetchRooms = func(name, roomType string) {
			room, err = testCon.CreateRoom(name, roomType)

			So(err, ShouldBeNil)
			So(room.ID, ShouldNotBeEmpty)

			rooms, err = testCon.GetRooms()

			So(err, ShouldBeNil)

			status, err = testCon.DeleteRoom(room.ID)

			So(status, ShouldEqual, "success")
		}
	)

	Convey("Tests for GetRooms", t, func() {
		Convey("GetRooms should return a slice of Room values that is not empty", func() {
			fetchRooms("Brookshire", "livingspace")

			So(err, ShouldBeNil)

			for _, newRoom := range rooms {
				So(newRoom.Capacity, ShouldEqual, 6)
				So(newRoom.ID, ShouldNotBeEmpty)
				So(newRoom.Name, ShouldEqual, "Brookshire")
				So(newRoom.Type, ShouldEqual, "livingspace")
			}
		})

		Convey("Print the successfully fetched Room", func() {
			fetchRooms("Mogadishu", "Office")

			So(err, ShouldBeNil)
			for _, newRoom := range rooms {
				printRoom(newRoom)
			}
		})
	})
}

// TestUpdateRoom tests the functionality of UpdateRoom
func TestUpdateRoom(t *testing.T) {
	var (
		err               error
		room              Room
		status, newStatus string

		fetchRoom = func(name, roomType string) {
			room, err = testCon.CreateRoom(name, roomType)

			So(err, ShouldBeNil)
			So(room.ID, ShouldNotBeEmpty)

			newStatus, err = testCon.UpdateRoom("New Name", room.ID)

			So(err, ShouldBeNil)

			status, err = testCon.DeleteRoom(room.ID)

			So(status, ShouldEqual, "success")
		}
	)

	Convey("Tests for UpdateRoom", t, func() {
		Convey("UpdateRoom should return status equal to empty", func() {
			fetchRoom("Yorkshire", "livingspace")

			So(err, ShouldBeNil)
			So(newStatus, ShouldEqual, "success")
		})

		Convey("Print the successfully fetched Status", func() {
			fetchRoom("Mogadishu", "livingspace")

			So(err, ShouldBeNil)

			log.Println()
			log.Println("Status :", newStatus)
			log.Println()
		})
	})
}
