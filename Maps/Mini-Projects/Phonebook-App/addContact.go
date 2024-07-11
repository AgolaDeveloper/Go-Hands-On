// this adds contact(s)
package main

import "fmt"

//function that adds contact to our Phonebook
//this func takes another function (that returns a map data struct) as a return value
func AddContact(phoneBook map[string]string) {

	//store return value of phoneBook Map data struct to phoneBookMap
	phoneBookMap := phoneBook

	//it only adds Last Name and Phone Number
	var name string
	var phone string

	fmt.Println("Adding New Contact to Your PhoneBook")

	fmt.Println("NAME: ")
	fmt.Scan(&name)

	fmt.Println("Phone Number: ")
	fmt.Scan(&phone)

	//then we write the contact details to our data structure
	phoneBookMap[name] = phone
	//it doesn't return anything; since it adds data to our Map Data Struct
}
