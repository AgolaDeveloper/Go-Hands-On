// this program file enables user to add item to the inventory
package main

import (
	"fmt"
)

// the function takes function from inventory file as a parameter
func AddItem(inventory map[string]map[int]int) {
	//initialize the parameter
	ourMap := inventory

	//prompt item from user
	//...and add to the inventory
	var itemName string
	var itemPrice int
	var itemQuantity int

	//fmt.Println("Enter Item>>")
	fmt.Println("Item Name>> ")
	fmt.Scan(&itemName)

	//[make] inner map
	innerInventory := make(map[int]int)

	fmt.Println("Price (per item)>> ")
	fmt.Scan(&itemPrice)

	fmt.Println("Quantity>> ")
	fmt.Scan(&itemQuantity)

	//check if inventory empty or not
	if len(ourMap) == 0 {
		//if empty go ahead and add item to the inventory

		//FIRST, populate inner map...price as key and quantity as value
		innerInventory[itemPrice] = itemQuantity

		//then assigne innerMap, as value, to outerMap
		//outerMap has itemName as its key

		ourMap[itemName] = innerInventory

	} else {
		//else if Inventory not empty
		//FIRST, check if itemName already exist
		_, itemAlreadyExist := ourMap[itemName]

		if itemAlreadyExist {
			//if it's TRUE; item already exist...
			//...then FIRST check if itemPrice is same

			//you range the innerMap associated with the itemName

			for k, val := range ourMap[itemName] {
				if k == itemPrice {
					val += itemQuantity

					//then update the inventory
					//first update innerMap
					innerInventory[k] = val

					//then update inventory
					ourMap[itemName] = innerInventory
				}

			}

		}
	}

}
