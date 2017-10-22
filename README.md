
### [Amity-Golang]

This is a Golang version of the original system built in python.

Visit [Amity Space Allocation System](http://andela-dmigwi.github.io/amity-space-allocation/) for more information.


Amity is an [Andela](http://andela.com) facility that has several rooms in it. A room can be
either a **Living Space** or an **Office Space**. An Office Space can accomodate a maximum of
6 people and the Living Space can accomodate a maximum of 4 at ago.  

The Amity Space Allocation allocate people either The Living Space or The Office Space.  

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

 - Create a `.env` file in the root folder with following configuration

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
go run amity.go


# Usage
 
## Create Room
*Command:* `create_room <room_name>...`  
 - Creates rooms in Amity. This command allows one to create room  
   **Single room**: `create_room Dojo` -> Create room called **Dojo**  
   **Multiple rooms**: `create_room Dojo,Krypton,Valhala` -> Creates three rooms: **Dojo**, **Krypton** and **Valhala**

 - After typing the create command, you will be prompted to type:  
   `O` for Office   
       or  
   `L` for Living Space  

 *This will be repeated for every room you create*  


## Add Person
 *Command:* `add_person <person_name> <FELLOW|STAFF> [wants_accommodation]`  
 - Adds a person to the system and allocates the person to a random room. wants_accommodation here is an optional argument which can be either ``Y`` or ``N``.  
The default value if it is not provided is `N`.  

## Reallocate Person
 *Command:* `reallocate_person <person_identifier> <new_room_name>`  
 - Reallocate the person with person_name to new_room_name.  

## Print Allocations
*Command:* `print_allocations [filename]`  
 - Prints a list of allocations onto the screen. The file name is optional, if its not provided, data is not printed in a file.  
  
    ``Room Name:  Narnia ``  
    ``-------------------------------------``  
    ``MEMBER 1, MEMBER 2, MEMBER 3``  

   
    ``Room Name:  Krypton``  
    ``-------------------------------------``  
    ``MEMBER 1, MEMBER 2``  


    ``Room Name:  Krypton``  
    ``-------------------------------------``  
    ``MEMBER 1, MEMBER 2``  
  
## Print Empty Rooms
*Command:* `print_empty_rooms [filename]`  
 - Prints a list of unallocated rooms to the screen. The file name is optional, if its not provided, data is not printed in a file.  

## Print Room
*Command:* `print_room <room_name>`  
 - Prints the names of all the people in ``room_name`` on the screen.  

## Print Unallocated
*Command:* `print_unallocated [filename]`  
 - Prints a list of unallocated people to the screen. The file name is optional, if its not provided, data is not printed in a file.  

**@Done By [Migwi-Ndung'u] (http://www.github.com/dmigwi)**  
