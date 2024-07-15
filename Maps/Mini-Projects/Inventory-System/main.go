package main

//entry point of the Application
func main() {
	//first, initialize inventory's method
	ourStruct := ourStruct()

	//add item to the inventory... passing inventory's func as argument
	AddItem(ourStruct)

	//check/display available items in the inventory...taking inventory's func as argument
	Items(ourStruct)
}
