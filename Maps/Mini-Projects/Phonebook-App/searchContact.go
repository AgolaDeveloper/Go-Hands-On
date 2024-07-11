// this searchs for a contact in our phonebook map data structure
package main

import "fmt"

//the func ranges through phonebook and displays contact details
func searchContact() {

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
				fmt.Printf("Name: %v\n Contact Number: %V\n", key, value)

			}
		}

	}
}
