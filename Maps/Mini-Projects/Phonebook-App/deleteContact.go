// this deletes a contact from our Phonebook [MAP Data Structure]
package main

import "fmt"

//this func deletes a contact from the phonebook
//must check if the contact is present though

//this func takes phoneBook data struct as a parameter and deletes a contact from it
func DeleteContact(phoneBook func() map[string]string) {
	//store return value of phonebook to phonBookMap
	phoneBookMap := phoneBook()

	//then checks if phonebook is empty first
	if phoneBookMap == nil {
		//if it's nill/empty >>then deletion impossible
		fmt.Println("Nothing to delete! PhoneBook Empty")
	} else {
		//else, phonebook get displayed and you choose contact to delete
		//by ranging through the entire phonebook Map

		fmt.Println("Your PhoneBook:")
		for key := range phoneBookMap {
			//we only display Names [keys] to the user to choose
			fmt.Println(key)
		}

		var contDelete string
		fmt.Println("Enter Contact to Delete:\n >>: ")
		fmt.Scan(&contDelete)

		//again check if contact entered doesn't exist in the data struct [Our Map in this case]
		_, ok := phoneBookMap[contDelete]

		if !ok {
			//if it doesn't exist
			fmt.Printf("\nError... %v Not in Phonebook\n", contDelete)
		} else {
			//else... if it exists delete the key from the Map [our phonebook data structure]
			delete(phoneBookMap, contDelete)

		}
	}

	//doesn't return any value
}
