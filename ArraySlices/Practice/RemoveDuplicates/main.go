// this is the main program file
// it's the entry point to our program
package main

func main() {
	//we first initialize our slice
	ourSlice := slice()
	//then populate it
	pop := popSlice(ourSlice)

	//pop := popSlice()
	//then remove duplicates
	removeDuplicate(pop)

}
