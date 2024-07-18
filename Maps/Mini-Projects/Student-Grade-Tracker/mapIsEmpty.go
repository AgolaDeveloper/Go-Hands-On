// this program file helps with checking whetehr our Map-Struct [that holds the Students' Grades is empty]
package main

//func that checks whether the map data structure is empty
//it takes our map data structure as a parameter and ...
//...returns boolean value [TRUE or FALSE]
func mapIsEmpty(studentGrade map[string]map[string]string) bool {
	//first initialize the map parameter to a local map variable
	studentGradeStruct := studentGrade

	//check whether empty

	if len(studentGradeStruct) == 0 {
		//if it's empty return true (implying...it's true ourMap is empty)
		return true
	} else {
		//else if map is not empty return false (implying... ourMap is not empty)
		return false
	}
}
