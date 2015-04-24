package driver

import ("fmt"
		"time"
		"strconv")

func Orders_make_state_matrix() [][]int {
	
	State_matrix := make([][]int, N_elevators)
	for i := range State_matrix {
		State_matrix[i] = make([]int, 3 + M_floors + (M_floors-1)*2)
		for j := range State_matrix[i] {
			State_matrix[i][j] = 0
		}
	}
	return State_matrix
}

func Orders_update_state_matrix(message Message, matrix [][]int, row int, value int) int {

	time.Sleep(200*time.Millisecond)	
	
	destination_floor,_		:= strconv.Atoi(message.Order_type[0:1])
	order_type				:= message.Order_type[1:len(message.Order_type)]
	
	if (order_type == "UP") && (destination_floor < M_floors) {
		start_i 		:= M_floors + 2
		end_i			:= start_i + destination_floor -1
		matrix[row][end_i] 	= value	
	} else if (order_type == "DOWN") && (destination_floor > 1) {
		start_i 		:= M_floors*2 + 1
		end_i			:= start_i + destination_floor - 2
		matrix[row][end_i] 	= value	
	} else if order_type == "ORDER" {
		start_i 		:= 2
		end_i			:= start_i + destination_floor - 1
		matrix[row][end_i] 	= value	
	} else {
		fmt.Println("Orders: invalid order_type")
		return 1
	}
	return 0
}
