// this program file enables user to add item to the inventory
package main

import (
	"fmt"
	"strings"
)

// the function takes function from inventory file as a parameter
func AddItem(inventory map[string]map[int]int) {
	//initialize the parameter
	ourMap := inventory

	//[make] inner map
	innerInventory := make(map[int]int)
	var itemName string
	var itemPrice int
	var itemQuantity int

	//check if inventory empty or not
	if len(ourMap) == 0 {
		//if empty go ahead and add item to the inventory
		//fmt.Println("Enter Item>>")
		fmt.Println("Item Name>> ")
		fmt.Scan(&itemName)

		fmt.Println("Price (per item)>> ")
		fmt.Scan(&itemPrice)

		fmt.Println("Quantity>> ")
		fmt.Scan(&itemQuantity)

		//FIRST, populate inner map...price as key and quantity as value
		innerInventory[itemPrice] = itemQuantity

		//then assigne innerMap, as value, to outerMap
		//outerMap has itemName as its key

		ourMap[itemName] = innerInventory

	} else {
		//else if Inventory not empty
		//FIRST, check if itemName already exist
		count := 0

		fmt.Println("Item Name>> ")
		fmt.Scan(&itemName)

		for key, value := range ourMap {
			//range our map and check if our itemName matches any item already in the inventory
			itemAlreadyExist := strings.EqualFold(itemName, key)

			if itemAlreadyExist {
				count++
				//if it's TRUE; item already exist...
				//...then FIRST check if itemPrice is same
				fmt.Println("Price (per item)>> ")
				fmt.Scan(&itemPrice)

				//you range the innerMap associated with the itemName

				for k, val := range value {
					if k == itemPrice {
						//if prices match... then enter its qnty

						fmt.Println("Quantity>> ")
						fmt.Scan(&itemQuantity)

						val += itemQuantity

						//then update the inventory
						//first update innerMap
						innerInventory[k] = val

						//then update inventory
						ourMap[key] = innerInventory
					} else {
						//else, if itemPrice isn't equal to price of Item already in the inventory
						//...just add it as a distinctively different item to the inventory

						fmt.Println("Quantity>> ")
						fmt.Scan(&itemQuantity)

						//FIRST, populate inner map...price as key and quantity as value
						innerInventory[itemPrice] = itemQuantity

						//then assigne innerMap, as value, to outerMap
						//outerMap has itemName as its key

						ourMap[itemName] = innerInventory
					}

				}

			}
		}

		if count == 0 {
			//if count==0; itemName doesn't exist already
			//...go ahead and add item to the inventory

			fmt.Println("Price (per item)>> ")
			fmt.Scan(&itemPrice)

			fmt.Println("Quantity>> ")
			fmt.Scan(&itemQuantity)

			//FIRST, populate inner map...price as key and quantity as value
			innerInventory[itemPrice] = itemQuantity

			//then assigne innerMap, as value, to outerMap
			//outerMap has itemName as its key

			ourMap[itemName] = innerInventory
		}

	}

}
