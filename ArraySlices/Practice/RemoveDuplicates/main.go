// this is the main program file
// it's the entry point to our program
package main

func main() {
	//we first initialize our slice
	ourSlice := Slice()
	//then populate it
	pop := PopSlice(ourSlice)

	//pop := popSlice()
	//then remove duplicates
	remove := RemoveDuplicate(pop)

	DisplaySlices(pop, remove)
}
