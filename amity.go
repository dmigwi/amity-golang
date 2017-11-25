package main

import (
	"amity-golang/controllers"
	"fmt"
	"os"

	"github.com/urfave/cli"
)

// con defines Datastore intreface that is used to connect to the database
var con = controllers.GetConnection(controllers.DBConfig)

// addPerson adds a new user to the system
func addPerson(c *cli.Context) error {
	var (
		args      = c.Args()
		user, err = controllers.AddPerson(con, args.Get(0), args.Get(1), args.Get(2), args.Get(3), args.Get(4))
	)

	exitOnError(err)

	printUserDetails(user)

	return nil
}

// createRoom create a new room or returns it if exists
func createRoom(c *cli.Context) error {
	var (
		args      = c.Args()
		room, err = controllers.CreateRoom(con, args.Get(0), args.Get(1))
	)

	exitOnError(err)

	printRoomDetails(room)

	return nil
}

// exitOnError if err is passed the the program exists with an error message
func exitOnError(err error) {
	if err != nil {
		fmt.Println("Operation failed : ", err.Error())
		os.Exit(1)
	}
}

// printAllocations prints all rooms with occupants
func printAllocations(c *cli.Context) error {
	var rooms, err = controllers.GetRoomAllocations(con)

	exitOnError(err)

	for _, room := range rooms {
		printRoomDetails(room)
	}

	return nil
}

// printRoom prints details of the select room if it exists
func printRoom(c *cli.Context) error {
	var (
		args      = c.Args()
		room, err = controllers.GetRoom(con, args.Get(0))
	)

	exitOnError(err)

	printRoomDetails(room)

	return nil
}

// printRoomDetails prints all the underlying details of a room
func printRoomDetails(room controllers.Room) {
	fmt.Println()
	fmt.Println("Room Details")
	fmt.Println("=====================")
	fmt.Println("Capacity :", room.Capacity)
	fmt.Println("ID :", room.ID)
	fmt.Println("Name :", room.Name)
	fmt.Println("Type :", room.Type)
	fmt.Println("Occupants -")
	for index, user := range room.Occupants {
		fmt.Println("\tCount :", index+1)
		fmt.Println("\tFirstName :", user.FirstName)
		fmt.Println("\tID :", user.ID)
		fmt.Println("\tLastName", user.LastName)
		fmt.Println("\tType :", user.Type)
		fmt.Println()
	}
	fmt.Println("=====================")
	fmt.Println()

}

// printUserDetails prints the underlying details of a User
func printUserDetails(user controllers.User) {
	fmt.Println()
	fmt.Println("User Details")
	fmt.Println("=====================")
	fmt.Println("FirstName :", user.FirstName)
	fmt.Println("ID :", user.ID)
	fmt.Println("LastName :", user.LastName)
	fmt.Println("UserType :", user.Type)
	fmt.Println("LivingSpace - ")
	fmt.Println("\tCapacity :", user.LivingSpace.Capacity)
	fmt.Println("\tID :", user.LivingSpace.ID)
	fmt.Println("\tName :", user.LivingSpace.Name)
	fmt.Println("\tType :", user.LivingSpace.Type)
	fmt.Println("Office - ")
	fmt.Println("\tCapacity :", user.Office.Capacity)
	fmt.Println("\tID :", user.Office.ID)
	fmt.Println("\tName :", user.Office.Name)
	fmt.Println("\tType :", user.Office.Type)
	fmt.Println("=====================")
	fmt.Println()
}

// printUnallocatedPeople prints all people without a livingspace or an office
func printUnallocatedPeople(c *cli.Context) error {
	var users, err = controllers.GetUnallocatedPeople(con)

	exitOnError(err)

	for _, user := range users {
		printUserDetails(user)
	}

	return nil
}

// reallocatePerson moves a user from one room to another
func reallocatePerson(c *cli.Context) error {
	var (
		args      = c.Args()
		room, err = controllers.ReallocatePerson(con, args.Get(0), args.Get(1), args.Get(2))
	)

	exitOnError(err)

	printRoomDetails(room)

	return nil
}

// getCommands sets up all the necessary commands
func getCommands() []cli.Command {

	return []cli.Command{
		{
			Name:    "create_room",
			Aliases: []string{"cr"},
			Usage:   "amity-golang create_room <room_name> <office|Livingspace>",
			Action:  createRoom,
		},
		{
			Name:    "add_person",
			Aliases: []string{"ap"},
			Usage:   "amity-golang add_person <First_Name> <Last_Name> <fellow|staff> <office_Name> <livingSpace_Name>",
			Action:  addPerson,
		},
		{
			Name:    "print_allocations",
			Aliases: []string{"pa"},
			Usage:   "amity-golang print_allocations",
			Action:  printAllocations,
		},
		{
			Name:    "print_room",
			Aliases: []string{"pr"},
			Usage:   "amity-golang print_room <room_name>",
			Action:  printRoom,
		},
		{
			Name:    "print_unallocated_people",
			Aliases: []string{"pu"},
			Usage:   "amity-golang print_unallocated_people",
			Action:  printUnallocatedPeople,
		},
		{
			Name:    "reallocate_person",
			Aliases: []string{"rp"},
			Usage:   "amity-golang reallocate_person <First_Name> <Last_Name> <New_Room_Name>",
			Action:  reallocatePerson,
		},
	}
}

// main initiates the application
func main() {
	var app = cli.NewApp()

	app.Name = "Amity"
	app.Usage = "Room allocation app"
	app.Description = "A golang application written to help room allocation in amity "
	app.Version = "1.0.0"
	app.Commands = getCommands()

	app.Run(os.Args)
}
