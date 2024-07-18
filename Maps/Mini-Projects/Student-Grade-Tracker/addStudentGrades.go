// this program file helps with adding students Grade to the tracker-Struct
package main

import (
	"fmt"
	"strings"
)

//func adds student and respective grades per subject
//it takes return value of studentGrade method as parameter...
//...(basically, it takes data structure as parameter then writes to it)

// and returns nothing
func addStudentGrades(studentGrade map[string]map[string]string) {
	//first initialize the parameter for readiness of local usage
	studentGradeStruct := studentGrade

	//make an empty map; an innerMap that'll act as the outerMap's value
	InnerStudentGrade := make(map[string]string)

	var studentName string
	var studentSubject string
	var subjectGrade string

	//now prompt user's entry for students details
	fmt.Println("Enter student's NAME: ")
	fmt.Scan(&studentName)
	//must know number of subjects to be recorded for every student

	//we check if MAP-STRUCT is empty
	var numOfSubjects int

	if len(studentGradeStruct) == 0 {
		//if empty, go ahead and add studentGrade
		fmt.Printf("How many Subjects to be recorded for %v \n", studentName)
		fmt.Scan(&numOfSubjects)

		for i := 0; i < numOfSubjects; i++ {
			//enter subject and respective grade

			fmt.Printf("Enter %v's SUBJECT %v \n: \n", studentName, i+1)
			fmt.Scan(&studentSubject)

			fmt.Printf("Enter %v's GRADE\n: \n", studentSubject)
			fmt.Scan(&subjectGrade)

			//then map every subject to its grade in the inner map
			InnerStudentGrade[studentSubject] = subjectGrade
		}

		//then finally map innerMap [subjects with their respective grades] to the student's Name
		studentGradeStruct[studentName] = InnerStudentGrade
	} else {
		//if the MAP-STRUCT isn't empty
		//must first check if the Name already exist in the MAP-STRUCT
		//...by first ranging through the MAP-STRUCT; and confirming if the Name to be entered matches any key-name
		count := 0
		for key := range studentGradeStruct {
			nameExist := strings.EqualFold(key, studentName)

			if nameExist {
				count++
				//if the name existprint->

				continue
			}
		}

		if count == 0 {
			//go ahead and add studentGrade

			fmt.Printf("How many Subjects to be recorded for %v \n", studentName)
			fmt.Scan(&numOfSubjects)

			for i := 0; i < numOfSubjects; i++ {
				//enter subject and respective grade

				fmt.Printf("Enter %v's SUBJECT %v \n: \n", studentName, i+1)
				fmt.Scan(&studentSubject)

				fmt.Printf("Enter %v's GRADE\n: \n", studentSubject)
				fmt.Scan(&subjectGrade)

				//then map every subject to its grade in the inner map
				InnerStudentGrade[studentSubject] = subjectGrade
			}

			//then finally map innerMap [subjects with their respective grades] to the student's Name
			studentGradeStruct[studentName] = InnerStudentGrade

		} else {
			//else if count is >1
			fmt.Println("No Duplicates Please! Name already exist")

		}
	}

}
