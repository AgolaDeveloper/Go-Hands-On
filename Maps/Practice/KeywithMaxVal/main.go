// this program Finds Key with Maximum Value in a map
package main

import "fmt"

func keyWithMaxValue(ourMap func() map[string]int) {
	//initialize variable to store whatever [map] returned by our parameter
	theMap := ourMap()

	//initially make maxValue to 0
	maxValue := 0
	var maxKey string

	//iterate through the map
	for key, value := range theMap {
		//then compare every value to it
		if value > maxValue {
			//make value of this key maximum value if it's greater than the current maxValue

			maxValue = value
			//and assign it's key as key with maxVal
			maxKey = key
		}
	}

	//here's our map
	fmt.Printf("\nThe MAP:\n %v \n", theMap)

	//Print the key with maxValue

	fmt.Printf("\nKey with Maximum Value: %v \n", maxKey)
	fmt.Printf("\nIt's Value: %v \n", maxValue)

}

func main() {
	//Call keykeyWithMaxValue function with popMap passed as its argument
	keyWithMaxValue(popMap)

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
