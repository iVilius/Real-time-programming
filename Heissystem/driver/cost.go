package driver

//Kalle whos alive

import(
        "fmt"
        "time"
)



func Cost_main(message Message, order_port string, state_matrix [][]int, active_elevator_list []int, receive_ch chan Message) {

	cost_array 	:= Cost_calculate_cost_array(active_elevator_list, state_matrix, message)
	
	lucky_elevator 	:= Cost_return_lucky_elevator(cost_array, active_elevator_list)
	
	time.Sleep(2000*time.Millisecond)
	Cost_send_order(lucky_elevator, order_port, message, receive_ch)
}

func Cost_check_if_intern_order(message Message, active_elevators []int, temp_cost_list []int, state_matrix [][]int) ([]int, int) {

	if (message.Order_type == "1ORDER" || message.Order_type == "2ORDER" || message.Order_type == "3ORDER" || message.Order_type == "4ORDER") {
		for i := 0; i < len(active_elevators); i++ {
			if active_elevators[i] == message.Trunc_IP {
			
				Orders_update_state_matrix(message, state_matrix, i, 1)
				temp_cost_list[i] = -1
				return temp_cost_list, 1
			}
		}
	}
	return temp_cost_list, 0
}

func Cost_send_order(elevator_to_send int, port string, message Message, receive_ch chan Message) {
        if elevator_to_send == -1 {
        	fmt.Println("Cost: error in assigning order to elevator")
        	return
        }
        fmt.Println("---Order", message.Order_type, "has been assigned to elevator", elevator_to_send) 
        var order Message
        var sleep_time int 		= 10
        
        terminate_ch 			:= make(chan int, 10)
        
        order.ID                = ORDER_ASSIGN // cofnirm order
        
        order.Order_elevator_ID = elevator_to_send
        order.Order_type        = message.Order_type
        
        go UDP_broadcast("129.241.187.255:" + port, order, sleep_time, terminate_ch)
        
        for {
        	select {
        	case i := <- receive_ch:
        		if i.ID == ORDER_ASSIGN_ACK {
        			if i.Trunc_IP == elevator_to_send {
        				fmt.Println("Elevator", i.Trunc_IP, ". Thank you for taking the order")
        				terminate_ch <- 1
        				return
        			}
        		}
        	}
        }

}

func Cost_return_lucky_elevator(array []int, active_elevators []int) int{
        
        var temp_index int      = 0
        var min int             = 50000
        
        for i := range array {
                if array[i] < min {
                        min = array[i]
                        temp_index = i
                }                         
        }
        
        if active_elevators[temp_index] == 0 {
                fmt.Println("Cost: no elevator deserves the order")
                return -1               
        }
        
        return active_elevators[temp_index]
}

func Cost_delete_order(port string, message Message, state_matrix [][]int, active_elevators []int ) {
	
	for i := 0; i < len(active_elevators); i++ {
			if active_elevators[i] == message.Trunc_IP {
				Orders_update_state_matrix(message, state_matrix, i, 0)
				fmt.Println("Following order deleted: ", message.Order_type)
			}
	}
}

func Cost_calculate_cost_array(active_elevators []int, state_matrix [][]int, message Message) []int {
        
    temp_cost_list  		:= make([]int, len(active_elevators))
	
	temp_cost_list, flag 	:= Cost_check_if_intern_order(message, active_elevators, temp_cost_list, state_matrix)
	
	if flag == 1 {
		fmt.Println("---The order is intern")
	}
	
	return temp_cost_list
}




