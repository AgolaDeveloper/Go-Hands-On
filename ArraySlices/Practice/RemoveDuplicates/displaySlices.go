// this program file displays original slice[with duplicates] and sorted one [without duplicates]
package main

import "fmt"

//it takes two parameters
//the original slice with duplicates [return value of popSlice]
//...& sorted slice without duplicates [return value of removeDuplicate]
func displaySlices(originalSlice, sortedSlice []int) {
	//print both slices

	fmt.Println("Original Slice [With Duplicates]: ", originalSlice)
	for index, value := range originalSlice {
		fmt.Printf("index %v: %v\n", index, value)
	}
	fmt.Println("____________________________________________________")

	fmt.Println("New Slice [Without Duplicates]: ", sortedSlice)
	for index, value := range sortedSlice {
		fmt.Printf("index %v: %v\n", index, value)

	}
}
