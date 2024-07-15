// this program file gives worth of all items in the inventory
package main

import "fmt"

//func takes return value of inventory.go's method as a parameter
func inventoryTotal(inventory map[string]map[int]int) int {
	//first initialize the parameter into a local variable
	ourItems := inventory

	totalStock := 0

	//first check if the inventory [ourItems] is empty
	if len(ourItems) == 0 {
		fmt.Println("Ooops...! Inventory empty")
	} else {
		//range through ourItems[the inventory] and sum-up the prices of all items as totalStock

		//total for every item is its item per piece * its quantity in the inventory
		//price of items is inner map's key and quantity is value of inner map

		for _, value := range ourItems {
			for k, val := range value {
				totalStock += (k * val)
			}
		}
	}

	return totalStock
}
