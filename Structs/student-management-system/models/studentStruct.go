// package that helps with creating our Data Structure that'll store the student details
package main

func ouStruct() {
	//create the struct that will hold the student details
	//this stuct shall hold slice of object students (another struct)
	type Students struct {
		Students []Student
	}

	func (*s Students) setStudents(students []Student){
s.Students= students
	}

	
}
