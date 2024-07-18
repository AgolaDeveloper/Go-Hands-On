// this program file helps with deleting student's Grade from the tracker-Struct
package main

import "fmt"

//func takes 2 parameters
//our Map-struct and return value of mapIsEmpty's func
func deleteStudentGrade(studentGrade map[string]map[string]string, mapEmpty bool) {

	//first initialize the mapparameter to a local map variable
	studentGradeStruct := studentGrade

	//then perform deletion only if the map-Struct ain't empty
	if mapEmpty {
		//if it's true that it's empty print->
		fmt.Printf("\nOops!...No student Records Found \nCan't perform Any Deletion Operation\n")
	} else {
		//else if map-Struct isn't empty
		//go ahead and delete only what's availale in the Map
		//else... print(No match)

		var student2Delete string
		fmt.Println("Enter Name of Student to Delete: ")
		fmt.Scan(&student2Delete)

		_, studentExist := studentGradeStruct[student2Delete]

		if !studentExist {
			//if it's false that student2Delete exist ->
			fmt.Printf("\n%v can't be found : No match!\n", student2Delete)

		} else {
			//else if there's a match->
			//(Delete student2Delete from the Map-struct)
			delete(studentGradeStruct, student2Delete)
		}
	}
}
