package driver

//Kalle whos alive

import(
        "fmt"
)



func Cost_main() {
	
	Cost_calculate_cost()
	Cost_return_lucky_elevator()
	Cost_send_order()
}

func Cost_calculate_cost (active_elevators []int, state_matrix [][]int, new_order string, elevator_ID int) {
        
        temp_cost_list          := make([]int, len(IP_list))
	
	






}

func Cost_send_order(elevator_to_send int, port string, new_order string) {
        
        var order Message
        var sleep_time int 	= 100
        receive_ch		:= make(chan Message, 500)
        terminate_ch 		:= make(chan int, 10)
        
        order.ID                = 5 // new order
        order.Order_elevator_ID = elevator_to_send
        order.Order_type        = new_order
        
        go UDP_broadcast("129.241.187.255:" + port, msg, sleep_time, terminate_ch)
        go UDP_receive(port, receive_ch, sleep_time)
        
        for {
                select {
                case i := <- receive_ch:
                	if i.ID = 6 {// Acknowledge order
                		if i.Trunc_IP == elevator_to_send {
                			terminate_ch <- 1
                		}
                	}
                }
        }
}

func Cost_return_lucky_elevator(array []int, active_elevators []int) int{
        
        var temp_index int      = -1
        var min int             = 50000
        
        for i := range array {
                if array[i] < min {
                        min = array[i]
                        temp_index = i
                }                         
        }
        
        if temp_index == -1 || active_elevators[temp_index] == 0 {
                fmt.Println("Cost: no elevator deserves the order")
                return -1                
        }
        return active_elevators[temp_index]
}
