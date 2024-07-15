package main

import "fmt"

//welcome func to the inventory management system
func welkam() {
	fmt.Println("INVENTORY ")
}

//entry point of the Application
func main() {
	//first, initialize inventory's method
	ourStruct := ourStruct()

retrry:
	var userChoice int

	fmt.Println("1. Stock Up")
	fmt.Println("2. Check Stock")

	fmt.Scan(&userChoice)

	switch userChoice {
	case 1:
		//add item to the inventory... passing inventory's func as argument
		fmt.Println("Stock Up! Add Item")
		AddItem(ourStruct)

	case 2:
		//check/display available items in the inventory...taking inventory's func as argument
		fmt.Println("Check ")
		Items(ourStruct)
	default:
		fmt.Println("Invalid Choice... Retry")
		goto retrry
	}

}
