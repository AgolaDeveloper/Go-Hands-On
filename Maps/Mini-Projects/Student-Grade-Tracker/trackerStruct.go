//this program file holds our Data Struct [Map in this case]
//it's the data struct in this file that'll store (in an organized way) our students' grades
//making it easier to perform different operations on the already organized data

package main

//defining func that creates our Data structure... the struct that'll store Students' grades
//it returns a map of studentName as key, its value an innerMap of key subjects and value grades

func studentGrades() map[string]map[string]string {
	//make an empty, but not nil, map

	studentGrades := make(map[string]map[string]string)

	return studentGrades
}
