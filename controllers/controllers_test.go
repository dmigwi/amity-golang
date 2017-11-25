package controllers

import (
	"amity-golang/models"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var dbCon *models.Connection

// TestAddPerson tests the functionality of AddPerson
func TestAddPerson(t *testing.T) {
	var (
		err            error
		user, testUser User

		addPerson = func(fname, lname, jobType, officeName, livingspaceName string) {
			if livingspaceName != "" {
				_, err = CreateRoom(con, livingspaceName, "livingspace")
			}

			if officeName != "" {
				_, err = CreateRoom(con, officeName, "office")
			}

			So(err, ShouldBeNil)

			user, err = AddPerson(con, fname, lname, jobType, officeName, livingspaceName)
		}
	)

	Convey("Tests for AddPerson", t, func() {
		Convey("AddPerson should return a User with all correct values present and an error", func() {
			var runValuesTests = func(rm models.Room) {
				So(rm.Capacity, ShouldNotEqual, 0)
				So(rm.ID, ShouldNotBeBlank)
				So(rm.Name, ShouldNotBeBlank)
				So(rm.Type, ShouldNotBeBlank)
			}

			addPerson("Jade", "Flora", "fellow", "Narnia", "PHP")

			So(err, ShouldBeNil)
			So(user.FirstName, ShouldNotBeBlank)
			So(user.ID, ShouldNotBeBlank)
			So(user.LastName, ShouldNotBeBlank)
			So(user.Type, ShouldNotBeBlank)
			So(user.User, ShouldNotResemble, testUser.User)

			runValuesTests(user.LivingSpace)
			runValuesTests(user.Office)
		})

		Convey("Print the successfully created User and Room ", func() {
			addPerson("Jackson", "Davids", "fellow", "", "Mogadishu")

			So(err, ShouldBeNil)

			printRooms()

			printUsers()

			So(con.DestroyData(), ShouldBeNil)
		})
	})
}

// TestCreateRoom tests the functionality of CreateRoom
func TestCreateRoom(t *testing.T) {
	var (
		err  error
		room Room

		createRoom = func(name, roomType string) {
			room, err = CreateRoom(con, name, roomType)
		}
	)

	Convey("Tests for CreateRoom ", t, func() {
		Convey("CreateRoom should return Room with all the correct values present and an error", func() {
			createRoom("Valhalla", "office")

			So(err, ShouldBeNil)
			So(room.Capacity, ShouldEqual, 6)
			So(room.ID, ShouldNotBeBlank)
			So(room.Name, ShouldNotBeBlank)
			So(room.Occupants, ShouldBeEmpty)
			So(room.Type, ShouldNotBeBlank)
		})

		Convey("Print the successfully created Room and User", func() {
			createRoom("Golang", "livingspace")

			So(err, ShouldBeNil)

			printRooms()

			printUsers()

			So(con.DestroyData(), ShouldBeNil)
		})
	})
}

// TestGetRoom tests the functionality of GetRoom
func TestGetRoom(t *testing.T) {
	var (
		err  error
		user User
		room Room

		fetchRoom = func(fname, lname, userType, rmName, rmType string) {
			room, err = CreateRoom(con, rmName, rmType)

			So(err, ShouldBeNil)
			So(room.ID, ShouldNotBeBlank)

			user, err = AddPerson(con, fname, lname, userType, room.Name, "")

			So(err, ShouldBeNil)
			So(user, ShouldNotResemble, (User{}))

			room = Room{}

			room, err = GetRoom(con, rmName)
		}
	)

	Convey("Tests for GetRoom ", t, func() {
		Convey("GetRoom should return Room with all the correct values present and an error", func() {
			fetchRoom("Joshua", "Mwaniki", "staff", "Valhalla", "office")

			So(err, ShouldBeNil)
			So(room.Capacity, ShouldEqual, 6)
			So(room.ID, ShouldNotBeBlank)
			So(room.Name, ShouldEqual, "Valhalla")
			So(room.Occupants, ShouldNotBeEmpty)
			So(room.Type, ShouldNotBeBlank)

			So(room.Occupants[0].FirstName, ShouldEqual, "Joshua")
			So(room.Occupants[0].ID, ShouldNotBeBlank)
			So(room.Occupants[0].LastName, ShouldEqual, "Mwaniki")
			So(room.Occupants[0].Type, ShouldEqual, "staff")
		})

		Convey("Print the successfully created Room and User", func() {
			fetchRoom("James", "Kaberia", "staff", "Pretoria", "office")

			So(err, ShouldBeNil)

			printRooms()

			printUsers()

			So(con.DestroyData(), ShouldBeNil)
		})
	})
}

// TestGetRoomAllocations tests the functionality of GetRoomAllocations
func TestGetRoomAllocations(t *testing.T) {
	var (
		err   error
		user  User
		room  Room
		rooms []Room

		fetchRoom = func(fname, lname, userType, rmName, rmType string) {
			room, err = CreateRoom(con, rmName, rmType)

			So(err, ShouldBeNil)
			So(room.ID, ShouldNotBeBlank)

			user, err = AddPerson(con, fname, lname, userType, room.Name, "")

			So(err, ShouldBeNil)
			So(user, ShouldNotResemble, (User{}))

			rooms, err = GetRoomAllocations(con)
		}
	)

	Convey("Tests for GetRoomAllocations ", t, func() {
		Convey("GetRoomAllocations should return a slice of Room with all the correct values present and an error", func() {
			fetchRoom("Garbrielle", "Wanjigi", "fellow", "Timbuktu", "office")

			So(err, ShouldBeNil)

			for _, room = range rooms {

				So(room.Capacity, ShouldEqual, 6)
				So(room.ID, ShouldNotBeBlank)
				So(room.Name, ShouldEqual, "Timbuktu")
				So(room.Occupants, ShouldNotBeEmpty)
				So(room.Type, ShouldNotBeBlank)

				So(room.Occupants[0].FirstName, ShouldEqual, "Garbrielle")
				So(room.Occupants[0].ID, ShouldNotBeBlank)
				So(room.Occupants[0].LastName, ShouldEqual, "Wanjigi")
				So(room.Occupants[0].Type, ShouldEqual, "fellow")
			}
		})

		Convey("Print the successfully created Room and User", func() {
			fetchRoom("Ashley", "Wanjigi", "fellow", "Golang", "office")

			So(err, ShouldBeNil)

			printRooms()

			printUsers()

			So(con.DestroyData(), ShouldBeNil)
		})
	})
}

// TestGetUnallocatedPeople tests the functionality of GetUnallocatedPeople
func TestGetUnallocatedPeople(t *testing.T) {
	var (
		err   error
		user  User
		users []User

		fetchUsers = func(fname, lname, userType string) {
			user, err = AddPerson(con, fname, lname, userType, "", "")

			So(err, ShouldBeNil)
			So(user, ShouldNotResemble, (User{}))

			users, err = GetUnallocatedPeople(con)
		}
	)

	Convey("Tests for GetUnallocatedPeople ", t, func() {
		Convey("GetUnallocatedPeople should return a slice of User with all the correct values present and an error", func() {
			fetchUsers("David", "Holmes", "staff")

			So(err, ShouldBeNil)

			for _, user = range users {
				So(user.FirstName, ShouldEqual, "David")
				So(user.ID, ShouldNotBeBlank)
				So(user.LastName, ShouldEqual, "Holmes")
				So(user.LivingSpace, ShouldResemble, (User{}))
				So(user.Office, ShouldResemble, (User{}))
				So(user.Type, ShouldEqual, "staff")
			}
		})

		Convey("Print the successfully created Room and User", func() {
			fetchUsers("Mark", "Oyaboade", "fellow")

			So(err, ShouldBeNil)

			printRooms()

			printUsers()

			So(con.DestroyData(), ShouldBeNil)
		})
	})
}

// TestReallocatePerson tests the functionality of ReallocatePerson
func TestReallocatePerson(t *testing.T) {
	var (
		err  error
		user User
		room Room

		fetchRoom = func(fname, lname, userType, oldRmName, newRmName, rmType string) {
			room, err = CreateRoom(con, oldRmName, rmType)

			So(err, ShouldBeNil)
			So(room.ID, ShouldNotBeBlank)

			user, err = AddPerson(con, fname, lname, userType, room.Name, "")

			So(err, ShouldBeNil)
			So(user, ShouldNotResemble, (User{}))

			room, err = CreateRoom(con, newRmName, rmType)

			So(err, ShouldBeNil)
			So(room.ID, ShouldNotBeBlank)

			room = Room{}

			room, err = ReallocatePerson(con, fname, lname, newRmName)
		}
	)

	Convey("Tests for ReallocatePerson ", t, func() {
		Convey("ReallocatePerson should return Room with all the correct values present and an error", func() {
			fetchRoom("James", "Bond", "fellow", "Valhalla", "Mogadishu", "office")

			So(err, ShouldBeNil)
			So(room.Capacity, ShouldEqual, 6)
			So(room.ID, ShouldNotBeBlank)
			So(room.Name, ShouldEqual, "Mogadishu")
			So(room.Occupants, ShouldNotBeEmpty)
			So(room.Type, ShouldNotBeBlank)

			So(room.Occupants[0].FirstName, ShouldEqual, "James")
			So(room.Occupants[0].ID, ShouldNotBeBlank)
			So(room.Occupants[0].LastName, ShouldEqual, "Bond")
			So(room.Occupants[0].Type, ShouldEqual, "fellow")
		})

		Convey("Print the successfully created Room and User", func() {
			fetchRoom("TheSecret", "Coder", "fellow", "Narnia", "Shell", "office")

			So(err, ShouldBeNil)

			printRooms()

			printUsers()

			// So(con.DestroyData(), ShouldBeNil)
		})
	})
}

// TestGetRoomDetails tests the functionality of getRoomDetails
func TestGetRoomDetails(t *testing.T) {
	var (
		err   error
		users []models.User
		rm    models.Room
		room  Room

		getRoomDetails = func(fname, lname, userType, rmName, rmType string) {
			room, err = CreateRoom(con, rmName, rmType)

			So(err, ShouldBeNil)
			So(room.ID, ShouldNotBeBlank)

			_, err = AddPerson(con, fname, lname, userType, room.Name, "")

			So(err, ShouldBeNil)

			rm, users, err = getRoomDetails(con, rmName, rmType)
		}
	)

	Convey("Tests for getRoomDetails ", t, func() {
		Convey("getRoomDetaiils should return the room, a slice of user and an error ", func() {
			getRoomDetails("Daniel", "Ikigai", "fellow", "Tsavoo", "office")

			So(err, ShouldBeNil)
			So(rm.Capacity, ShouldEqual, 6)
			So(rm.ID, ShouldNotBeBlank)
			So(rm.Name, ShouldEqual, "Tsavoo")
			So(rm.Type, ShouldEqual, "office")

			So(users[0].FirstName, ShouldEqual, "Daniel")
			So(users[0].ID, ShouldNotBeBlank)
			So(users[0].LastName, ShouldEqual, "Ikigai")
			So(users[0].Type, ShouldEqual, "fellow")
		})

		Convey("Print the successfully created Room and User", func() {
			getRoomDetails("Christine", "J", "fellow", "Tsavoo", "office")

			So(err, ShouldBeNil)

			printRooms()

			printUsers()

			So(con.DestroyData(), ShouldBeNil)
		})
	})
}
