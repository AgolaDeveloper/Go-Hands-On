// this program Finds Key with Maximum Value in a map
package main

import "fmt"

func main() {

}

//a func that helps with creating and populating a map
func popMap() map[string]int {
	//create empty map
	ourMap := make(map[string]int, 0)

	var dishes int
	//user populating the map
	fmt.Println("How many dishes do you wanna stock in?: ")
	fmt.Scan(&dishes)

	var dish string
	var price int
	for i := 0; i < dishes; i++ {
		//user enter name of the dish...
		fmt.Println("Enter Dish Name: ")
		fmt.Scan(&dish)

		//...and its respective price
		fmt.Printf("\nEnter %v's Price: \n", dish)
		fmt.Scan(&price)

		//then map dish to its respective price [and stored as key-value pairin our map]
		ourMap[dish] = price
	}

	return ourMap
}
