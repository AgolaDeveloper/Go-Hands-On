// this file program helps with populating our basic data structure[slice]
// it prompts users to enter data elements into the slice for storage
package main

import "fmt"

//this function takes return value of ourSlice function [the Slice data struct]
//then returns the already populated slice

func popSlice(sliceStruct []int) []int {

	//store our parameter to a variable
	ourSlice := sliceStruct

	var element int

	var sliceElem int
	fmt.Println("How many elements do you wanna add to the Slice?")
	fmt.Scan(&sliceElem)

	for i := 0; i < sliceElem; i++ {
		//prompt user to enter data-element into the slice
		fmt.Println("Enter Element to our Slice: ")
		fmt.Scan(&element)

		//then append the element into the slice
		ourSlice = append(ourSlice, element)
	}

	return ourSlice
}
