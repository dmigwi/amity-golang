package models

import (
	"log"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// printUser prints to the console the underlying values of a User
var printUser = func(user UserSpaces) {
	log.Println()
	log.Println("FirstName :", user.FirstName)
	log.Println("ID :", user.ID)
	log.Println("LastName :", user.LastName)
	log.Println("LivingSpaceID :", user.LivingSpaceID)
	log.Println("OfficeID :", user.OfficeID)
	log.Println("Type :", user.Type)
	log.Println()
}

// TestCreateUsers tests the functionality of CreateUser
func TestCreateUsers(t *testing.T) {
	var (
		err    error
		status string
		user   UserSpaces

		fetchUser = func(fname, lname, userType, officeID, livingspaceID string) {
			user, err = testCon.CreateUser(fname, lname, userType, officeID, livingspaceID)

			So(err, ShouldBeNil)

			status, err = testCon.DeleteUser(user.ID)

			So(status, ShouldEqual, "success")
		}
	)

	Convey("Tests for CreateUser Method", t, func() {
		Convey("CreateUser should return User values that are not empty and a nil error", func() {
			fetchUser("Ashley", "Mwenje", "fellow", "sample-ID", "")

			So(err, ShouldBeNil)
			So(user.FirstName, ShouldEqual, "Ashley")
			So(user.ID, ShouldNotBeBlank)
			So(user.LivingSpaceID, ShouldBeBlank)
			So(user.OfficeID, ShouldEqual, "sample-ID")
			So(user.LastName, ShouldEqual, "Mwenje")
			So(user.Type, ShouldEqual, "fellow")
		})

		Convey("Print the succesfully created User", func() {
			fetchUser("David", "Owour", "fellow", "tyteuyte-uw", "hjahdjad")

			So(err, ShouldBeNil)

			printUser(user)

		})
	})
}

// TestDeleteUser tests the functionality of DeleteUser
func TestDeleteUser(t *testing.T) {
	var (
		err    error
		user   UserSpaces
		status string

		deleteUser = func(fname, lname, userType, officeID, livingspaceID string) {
			user, err = testCon.CreateUser(fname, lname, userType, officeID, livingspaceID)

			So(err, ShouldBeNil)
			So(user.ID, ShouldNotBeEmpty)

			status, err = testCon.DeleteUser(user.ID)

		}
	)

	Convey("Tests for DeleteUser", t, func() {
		Convey("DeleteUser should return a status success and a nil error", func() {
			deleteUser("Kryptonite", "Batman", "staff", "Office", "")

			So(err, ShouldBeNil)
			So(status, ShouldEqual, "success")
		})

		Convey("Print successfully fetched status ", func() {
			deleteUser("Kryptonite", "Batman", "staff", "Office", "")

			log.Println()
			log.Println("Status :", status)
			log.Println()
		})
	})
}

// TestGetUser tests the functionality of GetUser
func TestGetUser(t *testing.T) {
	var (
		err           error
		user, newUser UserSpaces
		status        string

		fetchUser = func(fname, lname, userType, officeID, livingspaceID string) {
			user, err = testCon.CreateUser(fname, lname, userType, officeID, livingspaceID)

			So(err, ShouldBeNil)
			So(user.ID, ShouldNotBeEmpty)

			newUser, err = testCon.GetUser(fname, lname, "")

			So(err, ShouldBeNil)

			status, err = testCon.DeleteUser(user.ID)

			So(status, ShouldEqual, "success")
		}
	)

	Convey("Tests for GetUser", t, func() {
		Convey("GetUser should return User values that are not empty", func() {
			fetchUser("Adam", "Shire", "fellow", "", "livingspace")

			So(err, ShouldBeNil)
			So(newUser.FirstName, ShouldEqual, "Adam")
			So(newUser.ID, ShouldNotBeEmpty)
			So(newUser.LastName, ShouldEqual, "Shire")
			So(newUser.LivingSpaceID, ShouldEqual, "livingspace")
			So(newUser.OfficeID, ShouldBeEmpty)
			So(newUser.Type, ShouldEqual, "fellow")
		})

		Convey("Print the successfully fetched User", func() {
			fetchUser("Adam", "Lupu", "staff", "Narnia", "")

			So(err, ShouldBeNil)

			printUser(newUser)
		})
	})
}

// TestGetUsers tests the functionality of GetUsers
func TestGetUsers(t *testing.T) {
	var (
		err    error
		user   UserSpaces
		users  []User
		status string

		fetchUsers = func(fname, lname, userType, officeID, livingspaceID string) {
			user, err = testCon.CreateUser(fname, lname, userType, officeID, livingspaceID)

			So(err, ShouldBeNil)
			So(user.ID, ShouldNotBeEmpty)

			users, err = testCon.GetUsers(officeID, livingspaceID)

			So(err, ShouldBeNil)

			status, err = testCon.DeleteUser(user.ID)

			So(status, ShouldEqual, "success")
		}
	)

	Convey("Tests for GetUsers", t, func() {
		Convey("GetUsers should return a slice of User values that is not empty", func() {
			fetchUsers("Joshua", "Wafula", "fellow", "", "uyweu2736tza")

			So(err, ShouldBeNil)
			So(users, ShouldNotBeEmpty)

			for _, newUser := range users {
				So(newUser.FirstName, ShouldEqual, "Joshua")
				So(newUser.ID, ShouldNotBeEmpty)
				So(newUser.LastName, ShouldEqual, "Wafula")
				So(newUser.Type, ShouldEqual, "fellow")
			}
		})

		Convey("Print the successfully fetched User", func() {
			fetchUsers("Joshua", "Wafula", "staff", "Camelot", "")

			So(err, ShouldBeNil)
			So(users, ShouldNotBeEmpty)

			for _, newUser := range users {
				log.Println()
				log.Println("FirstName :", newUser.FirstName)
				log.Println("ID :", newUser.ID)
				log.Println("LastName :", newUser.LastName)
				log.Println("Type :", newUser.Type)
				log.Println()
			}
		})
	})
}

// TestUpdateUser tests the functionality of UpdateUser
func TestUpdateUser(t *testing.T) {
	var (
		err               error
		user              UserSpaces
		status, newStatus string

		fetchUser = func(fname, lname, userType, officeID, livingspaceID string) {
			user, err = testCon.CreateUser(fname, lname, userType, officeID, livingspaceID)

			So(err, ShouldBeNil)
			So(user.ID, ShouldNotBeEmpty)

			newStatus, err = testCon.UpdateUser("FirstName", "LastName", user.ID, "Kilimani", "Dojo")

			So(err, ShouldBeNil)

			status, err = testCon.DeleteUser(user.ID)

			So(status, ShouldEqual, "success")
		}
	)

	Convey("Tests for UpdateUser", t, func() {
		Convey("UpdateUser should return status equal to empty", func() {
			fetchUser("Ekuru", "Aukot", "fellow", "State House", "State House")

			So(err, ShouldBeNil)
			So(newStatus, ShouldEqual, "success")
		})

		Convey("Print the successfully fetched Status", func() {
			fetchUser("Ekuru", "Aukot", "fellow", "State House", "State House")

			So(err, ShouldBeNil)

			log.Println()
			log.Println("Status :", newStatus)
			log.Println()
		})
	})
}
