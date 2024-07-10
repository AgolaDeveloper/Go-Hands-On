// a function that counts the frequency of each element in a slice and stores the result in a map.
package main

import "fmt"

func cntFrequency(ourSlice func() []int) {
	//now we apply nested for loop to count frequency of every element
	elements := ourSlice()

	//we create a map for storing frequency of every element
	frequentMap := make(map[int]int)

	//we check frequency of every number by using a nested loop
	for i := 0; i < len(elements); i++ {
		count := 0

		//we only check frequency of a number if it's not already in the map
		//else we continue to the next number

		_, ok := frequentMap[elements[i]]

		if ok == false {
			//if its presence in the map is false; we go ahead and check its frequency
			//...and then increment our its count
			for j := i; j < len(elements); j++ {
				if elements[i] == elements[j] {
					count++
				}
			}
			//...then map it to its resultant count[its frequency]
			frequentMap[elements[i]] = count

		}

	}

	//now you print frequency of every element
	fmt.Println("Here's our Slice: ", elements)

	fmt.Println("Here's our Map: ", frequentMap)
	for key, value := range frequentMap {
		fmt.Printf("\n%v: appears %v times\n", key, value)
	}
}

func main() {
	//callin function that helps with creating and populating our slice
	cntFrequency(ourSlice)
}

//function that makes an empty slice and populate user input into it
func ourSlice() []int {
	//creating an empty slice of integers
	ourSlice := make([]int, 0)
	var num int
	var intNums int

	fmt.Println("How many integer numbers do you wanna add to your slice? :")
	fmt.Scan(&intNums)

	for i := 0; i < intNums; i++ {
		fmt.Printf("\n Enter: ")
		fmt.Scan(&num)
		ourSlice = append(ourSlice, num)
	}
	return ourSlice
}
