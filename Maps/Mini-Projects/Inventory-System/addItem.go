// this program file enables user to add item to the inventory
package main

import "fmt"

//the function takes function from inventory file as a parameter
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

	//first, check if itemName exists
	_, ok := ourMap[itemName]

	//if itemName already exist in the inventory
	if ok {
		//just update its price and quantity

		for key, value := range ourMap {
			//check if key is the itemName
			if key == itemName {

				for k, val := range value {

					//again check if itemPrice equals key of this innerMap too

					if k == itemPrice {
						//then update by summing itemQuantity to innerMap's val, the intial/original quantity

						val += itemQuantity

					} else {
						//if not...then it should be added to the inventory distinctively

						//populate inner map with price, as key,and quantity as value
						innerInventory[itemPrice] = itemQuantity

						//now populate the inventory
						//...with itemName askey and inneInventory as value
						ourMap[itemName] = innerInventory

					}
				}
			}
		}

	} else {
		//and if the itemName doesn't exist add it to the inventory distinctively
		//populate inner map with price, as key,and quantity as value
		innerInventory[itemPrice] = itemQuantity

		//now populate the inventory
		//...with itemName askey and inneInventory as value
		ourMap[itemName] = innerInventory
	}

}
