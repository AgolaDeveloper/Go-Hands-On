// this adds contact(s)
package main

import "fmt"

//function that adds contact to our Phonebook
func addContact() {
	//it only adds Last Name and Phone Number
	var name string
	var phone string

	fmt.Println("Adding New Contact to Your PhoneBook")

	fmt.Println("NAME: ")
	fmt.Scan(&name)

	fmt.Println("Phone Number: ")
	fmt.Scan(&phone)

	//then we add the Contact details to our Phonebook Data Structure; which is a MAP for now

}
