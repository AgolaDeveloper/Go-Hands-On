// this is the phonebook; displays everything in our phonebook Map struct
package main

import "fmt"

//func checks if phonebook is empty then display all the contact
//it takes phonebook Map func as an argument

func DisplayPhonebook(phonebook map[string]string) {

	phoneBookMap := phonebook

	//first checks whether phonebook is empty or not
	if len(phoneBookMap) == 0 {

		//if it's empty >>then print empty
		fmt.Println("PhoneBook Empty")
	} else {
		//else, phonebook get displayed
		//by ranging through the entire phonebook Map

		fmt.Println("Your PhoneBook:")
		fmt.Println("*______________________________________________*")
		fmt.Println()
		for key, value := range phoneBookMap {
			//we only display Names [keys] to the user to choose
			fmt.Printf("***\n%v:\n%v\n\n", key, value)
		}
		fmt.Println()
		fmt.Println("*______________________________________________*")

	}
}
