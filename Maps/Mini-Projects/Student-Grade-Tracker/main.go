// this is the main entry program to our mini project
package main

import "fmt"

//main entry point
func main() {
	//first initialize func of our Data Struct
	ourStruct := studentGrades()

	for {

	retry:
		var userChoice int

		fmt.Println("CHOOSE A FUNCTIONALITY:")
		fmt.Println(" 1. ADD STUDENT GRADES:")
		fmt.Println(" 2. VIEW STUDENT GRADES:")
		fmt.Println(" 3. DELETE STUDENT GRADES:")
		//fmt.Println("CHOOSE A FUNCTIONALITY:")

		fmt.Scan(&userChoice)

		switch userChoice {
		case 1:
			//addStudentGrade
			addStudentGrades(ourStruct)
		case 2:
			//viewStudentGrade
			viewStudentGrade(ourStruct, mapIsEmpty(ourStruct))
		case 3:
			//deleteStudentGrade
			deleteStudentGrade(ourStruct, mapIsEmpty(ourStruct))
		default:
			fmt.Println("Invalid Choice: Retry>>>")
			goto retry
		}

	retrry:
		var userContinue string
		fmt.Printf("\nWanna Continue? \n (y/n): \n")
		fmt.Scan(&userContinue)

		if userContinue == "y" || userContinue == "Y" {
			continue
		} else if userContinue == "n" || userContinue == "N" {
			fmt.Println("Cheers!")
			break
		} else {
			fmt.Println("Invalid Choice: Retry>>>")
			goto retrry
		}
	}
}
