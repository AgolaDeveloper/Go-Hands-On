// this program file helps with viewing student's Grade to the tracker-Struct
package main

import "fmt"

//this func takes ourStruct [map] as a parameter
//plus bool parameter that checks whethe map is empty or not

//...and prints/displays the studentsGrades

func viewStudentGrade(studentGrade map[string]map[string]string, mapEmpty bool) {
	//first initialize the parameter as a local map variable
	studentGradeStruct := studentGrade

	//first check if map is empty or not

	if mapEmpty {
		//if it's TRUE that map is empty...print sth
		fmt.Println("Ooops!... No student Records Found")
	} else {
		//else if it's false [map isn't empty]
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
