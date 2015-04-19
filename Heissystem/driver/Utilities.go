package driver

import "fmt"

func Utilities_bubble_sort(array []int) {
	
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array)-1; j++ {
			if array[j] > array[j+1] {
				temp := array[j+1]
				array[j+1] = array[j]
				array[j] = temp			
			}
		}	
	}
}

func Utilities_find_collumn_in_state_matrix(value int, array []int) (int) {
	
	for i := 0; i < len(array); i ++ {
		if value == array[i] {
			return i
		}	
	}
	fmt.Println("Couldn't find the right collumn")
	return -1
}
