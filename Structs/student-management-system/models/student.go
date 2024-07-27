// this is a file that models student objects
package main

func student() struct{} {
	//create a student struct

	type Student struct {
		firstName       string
		lastName        string
		admissionNumber string
		class           string
		age             int
		subjects        map[string]int
	}
	return Student
}
