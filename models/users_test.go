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
	log.Println("Livingspace :", user.Livingspace)
	log.Println("Office :", user.Office)
	log.Println("Type :", user.Type)
	log.Println()
}

// TestCreateUsers tests the functionality of CreateUser
func TestCreateUsers(t *testing.T) {
	var (
		err    error
		status string
		user   UserSpaces

		fetchUser = func(fname, lname, office, livingspace, userType string) {
			user, err = testCon.CreateUser(fname, lname, office, userType, livingspace)

			So(err, ShouldBeNil)

			status, err = testCon.DeleteUser(user.ID)

			So(status, ShouldEqual, "success")
		}
	)

	Convey("Tests for CreateUser Method", t, func() {
		Convey("CreateUser should return User values that are not empty and a nil error", func() {
			fetchUser("Ashley", "Mwenje", "sample-ID", "", "fellow")

			So(err, ShouldBeNil)
			So(user.FirstName, ShouldEqual, "Ashley")
			So(user.ID, ShouldNotBeBlank)
			So(user.Livingspace, ShouldBeBlank)
			So(user.Office, ShouldEqual, "sample-ID")
			So(user.LastName, ShouldEqual, "Mwenje")
			So(user.Type, ShouldEqual, "fellow")
		})

		Convey("Print the succesfully created User", func() {
			fetchUser("David", "Owour", "tyteuyte-uw", "hjahdjad", "fellow")

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

		deleteUser = func(fname, lname, office, livingspace, userType string) {
			user, err = testCon.CreateUser(fname, lname, office, userType, livingspace)

			So(err, ShouldBeNil)
			So(user.ID, ShouldNotBeEmpty)

			status, err = testCon.DeleteUser(user.ID)

		}
	)

	Convey("Tests for DeleteUser", t, func() {
		Convey("DeleteUser should return a status success and a nil error", func() {
			deleteUser("Kryptonite", "Batman", "Office", "", "staff")

			So(err, ShouldBeNil)
			So(status, ShouldEqual, "success")
		})

		Convey("Print successfully fetched status ", func() {
			deleteUser("Kryptonite", "Batman", "Office", "", "staff")

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

		fetchUser = func(fname, lname, office, livingspace, userType string) {
			user, err = testCon.CreateUser(fname, lname, office, userType, livingspace)

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
			fetchUser("Adam", "Shire", "", "livingspace", "fellow")

			So(err, ShouldBeNil)
			So(newUser.FirstName, ShouldEqual, "Adam")
			So(newUser.ID, ShouldNotBeEmpty)
			So(newUser.LastName, ShouldEqual, "Shire")
			So(newUser.Livingspace, ShouldEqual, "livingspace")
			So(newUser.Office, ShouldBeEmpty)
			So(newUser.Type, ShouldEqual, "fellow")
		})

		Convey("Print the successfully fetched User", func() {
			fetchUser("Adam", "Lupu", "Narnia", "", "staff")

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

		fetchUsers = func(fname, lname, office, livingspace, userType string) {
			user, err = testCon.CreateUser(fname, lname, office, userType, livingspace)

			So(err, ShouldBeNil)
			So(user.ID, ShouldNotBeEmpty)

			users, err = testCon.GetUsers(office, livingspace)

			So(err, ShouldBeNil)

			status, err = testCon.DeleteUser(user.ID)

			So(status, ShouldEqual, "success")
		}
	)

	Convey("Tests for GetUsers", t, func() {
		Convey("GetUsers should return a slice of User values that is not empty", func() {
			fetchUsers("Joshua", "Wafula", "Intercon", "", "staff")

			So(err, ShouldBeNil)
			So(users, ShouldNotBeEmpty)

			for _, newUser := range users {
				So(newUser.FirstName, ShouldEqual, "Joshua")
				So(newUser.ID, ShouldNotBeEmpty)
				So(newUser.LastName, ShouldEqual, "Wafula")
				So(newUser.Type, ShouldEqual, "staff")
			}
		})

		Convey("Print the successfully fetched User", func() {
			fetchUsers("Joshua", "Wafula", "Camelot", "", "staff")

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

		fetchUser = func(fname, lname, office, livingspace, userType string) {
			user, err = testCon.CreateUser(fname, lname, office, userType, livingspace)

			So(err, ShouldBeNil)
			So(user.ID, ShouldNotBeEmpty)

			newStatus, err = testCon.UpdateUser("FirstName", "LastName", user.ID)

			So(err, ShouldBeNil)

			status, err = testCon.DeleteUser(user.ID)

			So(status, ShouldEqual, "success")
		}
	)

	Convey("Tests for UpdateUser", t, func() {
		Convey("UpdateUser should return status equal to empty", func() {
			fetchUser("Ekuru", "Aukot", "State House", "State House", "fellow")

			So(err, ShouldBeNil)
			So(newStatus, ShouldEqual, "success")
		})

		Convey("Print the successfully fetched Status", func() {
			fetchUser("Ekuru", "Aukot", "State House", "State House", "fellow")

			So(err, ShouldBeNil)

			log.Println()
			log.Println("Status :", newStatus)
			log.Println()
		})
	})
}
