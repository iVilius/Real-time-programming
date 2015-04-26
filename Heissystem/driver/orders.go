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

func Orders_update_elevator_queue(queue []int, message Message, value int) {

	destination_floor,_		:= strconv.Atoi(message.Order_type[0:1])
	order_type 				:= message.Order_type[1:len(message.Order_type)
	
	if (order_type == "UP") && (destination_floor < M_floors) {
		start_i 		:= M_floors + 2
		end_i			:= start_i + destination_floor -1
		matrix[end_i] 	= value
	} else if (order_type == "DOWN") && (destination_floor > 1) {
		start_i 		:= M_floors*2 + 1
		end_i			:= start_i + destination_floor - 2
		matrix[end_i] 	= value
	} else if order_type == "ORDER" {
		start_i 		:= 2
		end_i			:= start_i + destination_floor - 1
		matrix[end_i] 	= value
	} else {
		fmt.Println("Orders_queue: invalid order_type")
		return 1
	}
	return 0	
}

func Orders_execute_orders(queue []int) {
	
	elevator_direction 		:= Motor_get_direction()
	elevator_current_floor 	:= Sensors_get_latest_floor()
	
	//for {
	if elevator_direction == 1 {
		for i := 1; i <= M_floors-elevator_current_floor; i++ {
			if queue[2+M_floors-1+elevator_current_floor+i-1] == 1 { //UP orders
				Motor_set_destination_floor(elevator_current_floor+i)
				for (Sensors_get_latest_floor() != elevator_current_floor+i) {
					time.Sleep(5*time.Millisecond)
				}
				Lamps_door()
				queue[2+M_floors-1+elevator_current_floor+i-1] = 0
				queue[2+elevator_current_floor+i-1] == 0
				return
			} else if queue[2+elevator_current_floor+i-1] == 1 {	// ORDER(Internal) orders
				Motor_set_destination_floor(elevator_current_floor+i)
				for (Sensors_get_latest_floor() != elevator_current_floor+i) {
					time.Sleep(5*time.Millisecond)
				}
				Lamps_door()
				queue[2+elevator_current_floor+i-1] == 0
			} // DOWN orders
		}
	} /*else if elevator_direction == 2 {
		for i := 0; i < elevator_current_floor-2; i++ {
		
			if queue[2+M_floors-i+elevator_current_floor] == 1 { //DOWN orders
				Motor_set_destination_floor(elevator_current_floor-i-1)
				for (Sensors_get_latest_floor() != elevator_current_floor-i-1) {
					time.Sleep(5*time.Millisecond)
				}
				Lamps_door()
				queue[2+M_floors-i+elevator_current_floor] = 0
				queue[2+M_floors-i+elevator_current_floor-2*(M_floor-1)] == 0
				return
				
			} else if queue[2+M_floors-i+elevator_current_floor-2*(M_floor-1)] == 1 {
				Motor_set_destination_floor(elevator_current_floor-i-1)
				for (Sensors_get_latest_floor() != elevator_current_floor-i-1) {
					time.Sleep(5*time.Millisecond)
				}
				Lamps_door()
				queue[2+M_floors-i+elevator_current_floor-2*(M_floor-1)] == 0
			}
		}
	}*/
	
	
	
	
	
	
	
	
	/*
	//If any orders at the current floor
	for i := 0; i < M_floors; i++ {
		if current_floor == i+1 {					
			if queue[i+2] == 1 {					// ORDER(Internal) orders
				Lamps_door()
				queue[i+2] = 0
			} else if i != M_floors-1 {				
				if queue[M_floors+2+i] == 1{		// UP orders
					Lamps_door()
					queue[M_floors+2+i] = 0
				}
			} else if i != 0 {						
				if queue[2+M_floors+M_floors-1+i-1] == 1 {// DOWN orders
					Lamps_door()
					queue[M_floors+2+i] = 0
				}
			}
		}
	}
	
	//If any orders above current floor
	for i := 0; i < M_floors-1; i++ {
		for j := 0; j < M_floors-i-1; j++ {
			if queue[3+j] == 1 {					// ORDER(Internal) orders
				
			} else if i != M_floors-1-1 
				if queue[2+M_floors+M_floors-1+j] == 1 {// DOWN orders
					
				}
			}
		}
		for j := 0; j < M_floors-i-2; j++ {			
			if queue[M_floors+2+j+1] == 1 {			// UP orders
				
			}
		}
			
	}
	
	//If any order under current floor
	for i := 1; i < M_floors; i++ {					
		for j := 0; j < i; j++ {
			if queue[2+j] == 1 {					// ORDER(Internal) orders
				
			} else if queue[2+M_floors+j] == 1 {	// UP orders
				
			}
		}
		for j := 0; j < i -1; j++ {
			if queue[2+M_floors+M_floors-1+j] == 1{	// DOWN orders
				
			}	
		} 
	}
	
	*/
}




















