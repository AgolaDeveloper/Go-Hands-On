// here we write a program that inverts a map (swap keys and values).
package main

import "fmt"

//function that receives a map and inverts it (swap keys and values)
//ourMap func() map[int]int as its argument
func invertMap(ourMap func() map[int]int) {
	//initialize variable that'll hold return value of the the passed parameter
	theMap := ourMap()

	fmt.Println("Original Map: ", theMap)
	//we'll range through the map and invert it through swapping

	for key, value := range theMap {
		//then swap every key with its value
		temp := key
		key = value
		value = temp
	}

	fmt.Println("Inverted Map: ", theMap)

}

func main() {
	invertMap(popMap)

}

//a func that helps with creating and populating a map
func popMap() map[int]int {
	//create empty map
	ourMap := make(map[int]int, 0)

	var ourKeys int
	//user populating the map
	fmt.Println("How many keys do you wanna store in the map?: ")
	fmt.Scan(&ourKeys)

	var ourKey int
	var ourValue int
	for i := 0; i < ourKeys; i++ {
		//user enter name of the dish...
		fmt.Println("Enter Integer Key: ")
		fmt.Scan(&ourKey)

		//...and its respective price
		fmt.Printf("\nEnter %v's Value: \n", ourKey)
		fmt.Scan(&ourValue)

		//then map dish to its respective price [and stored as key-value pairin our map]
		ourMap[ourKey] = ourValue
	}

	return ourMap
}
