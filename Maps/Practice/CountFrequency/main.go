// a function that counts the frequency of each element in a slice and stores the result in a map.
package main

import "fmt"

func cntFrequency( ourSlice()[]int)

func main(){

}

func ourSlice() []int{
	//creating an empty slice of integers
ourSlice:=make([]int,0)
var num int
var intNums int

fmt.Println("How many integer numbers do you wanna add to your slice? :")
fmt.Scan(&intNums)

for i:=0;i<intNums;i++{
	fmt.Printf("\n Enter: ")
	fmt.Scan(&num)
	ourSlice = append(ourSlice, num)
}
return ourSlice
}