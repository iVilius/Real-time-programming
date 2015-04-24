package driver

import ("fmt"
		"time")

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

func Orders_update_state_matrix(matrix_ch chan [][]int, matrix [][]int, row int, m_floors int, order_type string, value int) {

	time.Sleep(200*time.Millisecond)
	fmt.Println(value)
	
	if (order_type == "UP") && (floor < m_floors) {
		start_i 		:= m_floors + 2
		end_i			:= start_i + floor -1
		matrix[row][end_i] 	= value	
	} else if (order_type == "DOWN") && (floor > 1) {
		start_i 		:= m_floors*2 + 1
		end_i			:= start_i + floor - 2
		matrix[row][end_i] 	= value	
	} else if order_type == "ORDER" {
		start_i 		:= 2
		end_i			:= start_i + floor - 1
		matrix[row][end_i] 	= value	
	} else {
		fmt.Println("Orders: invalid order_type")
	matrix_ch <- matrix
	}
}
