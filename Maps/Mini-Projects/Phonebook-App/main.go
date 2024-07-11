// This program creates a simple phonebook application where you can add, delete, and search for contacts using a map.
package main

import "fmt"

func main() {
	//your phonebook
	displayPhonebook(phonebook)

	var choice string
	fmt.Println("1. PHONEBOOK")
	fmt.Println("2. ADD Contact")
	fmt.Println("3. DELETE Contact")
	fmt.Println("4. SEARCH Contact")

	fmt.Scan(&choice)

	//add Contact... call addContact with phonebook passed as argument
	//... it'll write to phonebook['s Map Data Struct]
	addContact(phonebook)

	//Delete Contact... call addContact with phonebook passed as argument
	//... it'll delete an existing contact from the phonebook['s Map Data Struct]
	deleteContact(phonebook)

	//Delete Contact... call addContact with phonebook passed as argument
	//... it'll delete an existing contact from the phonebook['s Map Data Struct]
	searchContact(phonebook)

}
