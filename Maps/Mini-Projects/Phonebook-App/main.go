// This program creates a simple phonebook application where you can add, delete, and search for contacts using a map.
package main

import "fmt"

func main() {
	//your phonebook
	DisplayPhonebook(Phonebook)

	var choice int
	fmt.Println("1. PHONEBOOK")
	fmt.Println("2. ADD Contact")
	fmt.Println("3. DELETE Contact")
	fmt.Println("4. SEARCH Contact")

	fmt.Scan(&choice)

	switch choice {
	case 1:
		//your phonebook
		DisplayPhonebook(Phonebook)
	case 2:
		//add Contact... call addContact with phonebook passed as argument
		//... it'll write to phonebook['s Map Data Struct]
		AddContact(Phonebook)
	case 3:
		//Delete Contact... call addContact with phonebook passed as argument
		//... it'll delete an existing contact from the phonebook['s Map Data Struct]
		DeleteContact(Phonebook)
	case 4:
		//Delete Contact... call addContact with phonebook passed as argument
		//... it'll delete an existing contact from the phonebook['s Map Data Struct]
		SearchContact(Phonebook)

	}

}
