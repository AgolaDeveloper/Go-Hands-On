// this program file helps with deleting student's Grade from the tracker-Struct
package main

import (
	"fmt"
	"strings"
)

// func takes 2 parameters and deletes studentDetails/Grade
// our Map-struct and return value of mapIsEmpty's func
func deleteStudentGrade(studentGrade map[string]map[string]string, mapEmpty bool) {

	//first initialize the mapparameter to a local map variable
	studentGradeStruct := studentGrade

	//then perform deletion only if the map-Struct ain't empty
	if !mapEmpty {
		//else if map-Struct isn't empty
		//go ahead and delete only what's availale in the Map
		//else... print(No match)

		var student2Delete string
		fmt.Println("Enter Name of Student to Delete: ")
		fmt.Scan(&student2Delete)

		//first range through the map and...
		for key := range studentGradeStruct {
			//...CHECK if such a name exist in the map
			studentExist := strings.EqualFold(key, student2Delete)

			if studentExist {
				//else if there's a match->
				//(Delete KEY [not student2Delete] from the Map-struct)
				delete(studentGradeStruct, key)
				fmt.Printf("\n%v Successfully deleted\n", key)
			} else {
				//if it's false that student2Delete exist ->
				continue

			}
		}

	} else {
		//if it's true that it's empty print->
		fmt.Printf("\nOops!...No student Records Found \nCan't perform Any Deletion Operation\n")
	}
	//_, studentExist := studentGradeStruct[student2Delete]

}
