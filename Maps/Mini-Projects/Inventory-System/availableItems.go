// this program file gives access to every item available in the inventory
package main

import "fmt"

//function displays whether there is an item in the inventory
//takes inventory assa parameter
func Items(inventory map[string]map[int]int) {
	//intialize the inventory's method
	ourItems := inventory

	//checks whether inventory is empty and displays its state

	if len(ourItems) == 0 {
		//if its len is zero[it's empty]
		fmt.Println("Empty Inventory... Stock up!")
	} else {
		fmt.Println()
		fmt.Println("ITEMS in STOCK:")

		//...then range through the inventory and display all the items
		for key, value := range ourItems {

			//innerInventory/map
			for k, val := range value {
				//print outter map's key and its respective inner map's key,k, and value,val
				fmt.Printf("%v %v %v\n", key, k, val)
			}
		}
	}

}
