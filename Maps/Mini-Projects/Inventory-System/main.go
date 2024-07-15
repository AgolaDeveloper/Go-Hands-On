package main

import "fmt"

//welcome func to the inventory management system
func welkam() {
	fmt.Println("INVENTORY Management System -by TeneTek ")
}

//entry point of the Application
func main() {
	//first, initialize inventory's method
	ourStruct := ourStruct()

	for {
	retrry:
		var userChoice int

		fmt.Println("1. Stock Up (Add Item)")
		fmt.Println("2. Check Stock (Items in the Inventory)")
		fmt.Println("3. Check stockTotal (Total price of items in inventory)")

		fmt.Scan(&userChoice)

		switch userChoice {
		case 1:
			//add item to the inventory... passing inventory's func as argument
			fmt.Println("Stock Up! Add Item")
			AddItem(ourStruct)

		case 2:
			//check/display available items in the inventory...taking inventory's func as argument
			fmt.Println("Check Items in the Inventory ")
			Items(ourStruct)
		case 3:
			//check/display the total price of all items in the inventory/stock

			//first initialize return value from stockTotalPrice.go's method
			totalStock := inventoryTotal(ourStruct)
			fmt.Printf("\nTotal Stock[items' Price]: Ksh. %v \n", totalStock)
		default:
			fmt.Println("Invalid Choice... Retry")
			goto retrry
		}

	retry:
		var cont string
		fmt.Println("wanna continue? (y/n): ")
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
