
### [Amity-Golang]

This is the Amity room allocations system built in golang.

Visit [Amity Space Allocation System](http://andela-dmigwi.github.io/amity-space-allocation/) for more information.


Amity is an [Andela](http://andela.com) facility that has several rooms in it. A room can be
either a **Living Space** or an **Office Space**. An Office Space can accomodate a maximum of
six people and the Living Space can accomodate a maximum of four people at ago.  

The Amity Space Allocation allocate people either Livingspace or Officespace.  

A room can be allocated **ONLY** to a staff or a fellow at Andela. Staff cannot be allocated any living spaces.
 Fellows have a choice to choose a living space or not. 

### Installation
Clone `git clone https://github.com/dmigwi/amity-golang.git`

### Setup the Environment

 - Create a postgres database
 ```sql
 CREATE USER amity WITH PASSWORD "12345";

 CREATE DATABASE amity WITH OWNER amity;

 GRANT ALL PRIVILEGES ON DATABASE amity To amity;
 ```

 - Install the dependencies

 ```bash 
 glide install
 ```

 - Create a `.env` file in the root folder with following configuration:

 ```
export AG_DATABASE=amity
export AG_USER=amity
export AG_PASSWORD=12345
```

 
### Run the Tests

```bash
Chmod +x run_test.sh
./run_test.sh
```

### Run the system 

- Set `PATH` to create in an application:

```bash
export PATH=$PATH:$GOPATH/n
```

- Install the binaries 

```bash
go install
```

# Usage
 
## Create Room
*Command:* `amity-golang create_room <room_name> <room_type>`  
 - Creates rooms in Amity. This command allows one to create a new room 

## Add Person
 *Command:* `amity-golang add_person <First_Name> <Last_Name> <fellow|staff> <office_Name> <livingSpace_Name>`  
 - Adds a person to the system and allocates the person to the room (s) if they are available.  

## Reallocate Person
 *Command:* `amity-golang reallocate_person <First_Name> <Last_Name> <New_Room_Name>`  
 - Reallocate the person with person_name to new_room_name.  

## Print Allocations
*Command:* `amity-golang print_allocations`  
 - Prints a list of allocations onto the screen. 
  
    ``Room Details``  
    ``=====================``  
    ``Capacity : 6``  
    ``ID : 31e6c28c-8e55-47b9-a786-803ff29b3c2c``  
    ``Name : Narnia``  
    ``Type : office``  
    ``Occupants -``  
    ``=====================``    

## Print Room
*Command:* `amity-golang print_room <room_name>`  
 - Prints the names of all the people in ``room_name`` on the screen.  

## Print Unallocated People
*Command:* `amity-golang print_unallocated_people`  
 - Prints a list of unallocated people to the screen. The file name is optional, if its not provided, data is not printed in a file.  

**@[Migwi-Ndung'u] (http://www.github.com/dmigwi) 2017**  
