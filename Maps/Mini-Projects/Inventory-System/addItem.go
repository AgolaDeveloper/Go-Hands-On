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

	//WE'LL DO THIS ONLY IF INVENTORY AINT EMPTY
	//if its length isn't zero

	if len(ourMap) != 0 {
		//then we confirm if itemName already exist
		for key, value := range ourMap {
			//check if key[itemName already in the inventory] is...
			//... same as new itemName user wanna add to inventory
			//..using package strings' EqualFold method to compare/check that

			sameItemName := strings.EqualFold(key, itemName)

			if sameItemName {
				//if it's true the itemNames match,...
				//... go ahead and check if prices match too-in the inner map

				for k, val := range value {

					//again check if itemPrice equals key of this innerMap too

					if k == itemPrice {
						//then update by summing itemQuantity to innerMap's val, the intial/original quantity

						val += itemQuantity
						innerInventory[k] = val
						//first update innermap

						ourMap[key] = innerInventory

					} else {
						//if not...then it should be added to the inventory distinctively

						//populate inner map with price, as key,and quantity as value
						innerInventory[itemPrice] = itemQuantity

						//now populate the inventory
						//...with itemName askey and inneInventory as value
						ourMap[itemName] = innerInventory

					}
				}
			} else {
				//else if it doesn't exist already in the inventory...
				//populate inner map with price, as key,and quantity as value
				innerInventory[itemPrice] = itemQuantity

				//now populate the inventory
				//...with itemName askey and inneInventory as value
				ourMap[itemName] = innerInventory
			}
		}
	} else {
		//if inventory is empty...
		//...we straight and populate accordingly
		//populate inner map with price, as key,and quantity as value
		innerInventory[itemPrice] = itemQuantity

		//now populate the inventory
		//...with itemName askey and inneInventory as value
		ourMap[itemName] = innerInventory
	}
}
