package driver

import "fmt"

func Orders_make_state_matrix(N_elevators int, M_floors int) [][]int {
	
	State_matrix := make([][]int, N_elevators)
	for i := range State_matrix {
		State_matrix[i] = make([]int, 3 + M_floors + (M_floors-1)*2)
		for j := range State_matrix[i] {
			State_matrix[i][j] = 0
		}
	}
	return State_matrix
}

func Orders_make_new_order(matrix [][]int, row int, m_floors, floor int, floor_type string, value int) {
	
	if (floor_type == "UP") {
		start_i 	:= m_floors + 3
		end_i		:= m_floors*2 + 1
	} else if floor_type == "DOWN" {
		start_i 	:= m_floors*2 + 2
		end_i		:= m_floors*3

	} else if floor_type == "ORDER" {
		start_i 	:= 3
		end_i		:= m_floors + 2
	} else {
		fmt.Println("Orders: invalid floor_type")
	}

}
