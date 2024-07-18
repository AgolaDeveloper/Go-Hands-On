// this program file helps with viewing student's Grade to the tracker-Struct
package main

import "fmt"

//this func takes ourStruct [map] as a parameter
//...and prints/displays the studentsGrades

func viewStudentGrade(studentGrade map[string]map[string]string) {
	//first initialize the parameter as a local map variable
	studentGradeStruct := studentGrade

	//first check if map is empty or not

	if len(studentGradeStruct) == 0 {
		//if empty print sth
		fmt.Println("Ooops!... No student Records Found")
	} else {
		//else if not empty..
		//range through the mapStruct and display everything
		for key, value := range studentGradeStruct {
			//print student's Name first [it's key of the outterMap]
			fmt.Println(key)

			//then range its respective innerMap for both subjects,k, and grades,val
			for k, val := range value {
				fmt.Printf("%v: %v\n", k, val)

			}
			fmt.Println("*_____________________________________________*")
		}

	}
}
