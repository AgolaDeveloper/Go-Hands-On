// this program file helps with adding students Grade to the tracker-Struct
package main

import "fmt"

//func adds student and respective grades per subject
//it takes return value of studentGrade method as parameter...
//...(basically, it takes data structure as parameter then writes to it)

//and returns nothing
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

	var numOfSubjects int
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
}
