// this is the main entry program to our mini project
package main

//main entry point
func main() {
	//first initialize func of our Data Struct
	ourStruct := studentGrades()

	//addStudentGrade
	addStudentGrades(ourStruct)

	//viewStudentGrade
	viewStudentGrade(ourStruct, mapIsEmpty(ourStruct))

	//deleteStudentGrade
	deleteStudentGrade(ourStruct, mapIsEmpty(ourStruct))

}
