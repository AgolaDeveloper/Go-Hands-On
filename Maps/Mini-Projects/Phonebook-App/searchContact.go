// this searchs for a contact in our phonebook map data structure
package main

import "fmt"

//the func ranges through phonebook and displays contact details
func SearchContact(phonebook map[string]string) {
	phoneBookMap := phonebook

	//first confirms whether map is empty
	if len(phoneBookMap) == 0 {
		//if it's nill/empty >>then deletion impossible
		fmt.Println("PhoneBook Empty")
	} else {
		var cont string

		fmt.Println("Search>>\n >>: ")
		fmt.Scan(&cont)
		//first check if the search contact exist
		_, ok := phoneBookMap[cont]

		if !ok {
			//if it doesn't exist
			fmt.Printf("\nError... %v Not in Phonebook\n", cont)
		} else {
			//else... if it exists then show the contactName and Number
			for key, value := range phoneBookMap {
				//print key and value where key equals cont
				if key == cont {
					fmt.Printf("\n Name: %v\n Contact Number: %v\n", key, value)

				}
			}

		}

	}
}
