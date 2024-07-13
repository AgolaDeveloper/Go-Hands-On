// this file program receives already-populated slice & removes all the duplicates innit
package main

//receives already-populated slice as a parameter...
//...then removes all the duplicates
func removeDuplicate(sliceStruct []int) []int {
	//store received value of the parameter to a local variable
	ourSlice := sliceStruct
	//then append it to a whole new slice
	slice := make([]int, 0)
	slice = append(slice, ourSlice...)

	//create a new slice that'll help with slicing duplicates
	sliceDuplicates := make([]int, 0)

	//range the slice in nested loops while weeding out the duplicates
	for i := 0; i < len(slice); i++ {

		for j := i + 1; j < len(slice); j++ {

			if slice[i] == slice[j] {
				//if... the slice out the duplicate
				slice1 := slice[:j]
				sliceDuplicates = append(sliceDuplicates, slice1...)

				slice2 := slice[j+1:]
				sliceDuplicates = append(sliceDuplicates, slice2...)

				//then make sliceDupplicates 'slice'
				slice = sliceDuplicates
			}
		}
	}

	return slice
}
