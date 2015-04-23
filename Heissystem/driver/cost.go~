package driver

//Kalle whos alive

import(
        "fmt"
)



func Cost_main(message Message, order_port string, state_matrix [][]int, active_elevator_list []int) {

	cost_array 	:= Cost_calculate_cost_array(active_elevator_list, state_matrix, message)
	
	lucky_elevator 	:= Cost_return_lucky_elevator(cost_array, active_elevator_list)
	Cost_send_order(lucky_elevator, order_port, message)
	
}

func Cost_calculate_cost_array(active_elevators []int, state_matrix [][]int, message Message) []int {
        
        temp_cost_list  := make([]int, len(active_elevators))
	
	temp_cost_list, flag 	:= Cost_check_if_intern_order(message, active_elevators, temp_cost_list)
	
	if flag == 1 {
		fmt.Println("The order is intern")
		return temp_cost_list
	}
	
	
	
	
	
	
	
	return temp_cost_list
}
func Cost_check_if_intern_order(message Message, active_elevators []int, temp_cost_list []int) ([]int, int) {

	if (message.Order_type == "1" || message.Order_type == "2" || message.Order_type == "3" || message.Order_type == "4") {
		for i := 0; i < len(active_elevators); i++ {
			if active_elevators[i] == message.Trunc_IP {
				temp_cost_list[i] = -1
				return temp_cost_list, 1
			}
		}
	}
	return temp_cost_list, 0
}

func Cost_send_order(elevator_to_send int, port string, message Message) {
        
        var order Message
        var sleep_time int 	= 100
        receive_ch		:= make(chan Message, 500)
        terminate_ch 		:= make(chan int, 10)
        
        order.ID                = 5 // cofnirm order
        order.Order_elevator_ID = elevator_to_send
        order.Order_type        = message.Order_type
        
        go UDP_broadcast("129.241.187.255:" + port, order, sleep_time, terminate_ch)
        go UDP_receive(port, receive_ch, sleep_time)
        
        for {
        	select {
        	case i := <- receive_ch:
        		if i.ID == 6 {
        			if i.Trunc_IP == elevator_to_send {
        				terminate_ch <- 1
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
        } else if temp_index == -10{
       		fmt.Println("Cost: order already in queue") 
       		return -1
        }
        return active_elevators[temp_index]
}
