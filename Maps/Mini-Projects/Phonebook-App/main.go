// This program creates a simple phonebook application where you can add, delete, and search for contacts using a map.
package main

import "fmt"

func main() {
	phonebook := Phonebook()
	//your phonebook
	DisplayPhonebook(phonebook)
	for {
		var choice int
		fmt.Println("1. PHONEBOOK")
		fmt.Println("2. ADD Contact")
		fmt.Println("3. DELETE Contact")
		fmt.Println("4. SEARCH Contact")

		fmt.Scan(&choice)

		switch choice {
		case 1:
			//your phonebook
			DisplayPhonebook(phonebook)
		case 2:
			//add Contact... call addContact with phonebook passed as argument
			//... it'll write to phonebook['s Map Data Struct]
			AddContact(phonebook)
		case 3:
			//Delete Contact... call addContact with phonebook passed as argument
			//... it'll delete an existing contact from the phonebook['s Map Data Struct]
			DeleteContact(phonebook)
		case 4:
			//Delete Contact... call addContact with phonebook passed as argument
			//... it'll delete an existing contact from the phonebook['s Map Data Struct]
			SearchContact(phonebook)

		}

		fmt.Println("*****************************************************")

	retry:
		var cont string
		fmt.Println("Continue-(y/n) >>: ")
		fmt.Scan(&cont)

		if cont == "y" || cont == "Y" {
			continue
		} else if cont == "n" || cont == "N" {
			break
		} else {
			fmt.Println("Invalid Choice...retry")
			goto retry
		}
	}

}
