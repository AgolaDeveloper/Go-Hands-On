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
	if mapEmpty {
		//if it's true that it's empty print->
		fmt.Printf("\nOops!...No student Records Found \nCan't perform Any Deletion Operation\n")
	} else {
		//else if map-Struct isn't empty
		//go ahead and delete only what's availale in the Map
		//else... print(No match)

		//if not empty then ...
		//...ONLY DELETE IF THE NAME EXIST

		var student2Delete string
		fmt.Println("Enter Name of Student to Delete: ")
		fmt.Scan(&student2Delete)

		count := 0

		//first range through the map and check if the name2BeDeleted exist

		for key := range studentGradeStruct {
			//...CHECK if such a name exist in the map
			studentExist := strings.EqualFold(key, student2Delete)

			if studentExist {

				count++
				/*
					}

					if count == 0 {*/

				//if count==0; student exist
				//delete
				//(Delete KEY [not student2Delete] from the Map-struct)
				delete(studentGradeStruct, key)
				fmt.Printf("\n%v Successfully deleted\n", key)
			}

		}

		if count == 0 {
			//else if count>0; name doesn't exist->
			fmt.Printf("\n%v doesn't exist in our Records\n", student2Delete)
		}
	}
	//_, studentExist := studentGradeStruct[student2Delete]

}
