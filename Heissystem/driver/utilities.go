package driver

import ("fmt"
	"time"
)

func Utilities_bubble_sort_asc(array []int) {
	
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

func Utilities_find_column_in_state_matrix(value int, array []int) (int) {

	for i := 0; i < len(array); i ++ {
		if value == array[i] {
			return i
		}	
	}
	fmt.Println("Utilities: no column to use in state matrix")
	return -1
}

func Utilities_send_i_am_alive(port string) {
	
	var msg Message
	terminate_ch 			:= make(chan int, 1)
	
	msg.ID 					= 1

	go UDP_broadcast("129.241.187.255:" + port, msg, 10, terminate_ch)
}

func Utilities_broadcast_state(port string) {
	
	var msg Message
	msg.ID 					= 2
	terminate_ch 			:= make(chan int, 1)
	
	msg.Latest_floor 		= Sensors_get_latest_floor()
	msg.Direction 			= Motor_get_direction()

	go UDP_broadcast("129.241.187.255:" + port, msg, 10, terminate_ch)
	
	time.Sleep(100*time.Millisecond)
	terminate_ch 			<- 1
	
}

func Utilities_whos_alive(port string, IP_list []int, active_elevator_list_ch chan []int){
	
	listen_ch 				:= make(chan Message, 1024)
	auxilary_list			:= make([]int, len(IP_list))
	active_elevator_list	:= make([]int, len(IP_list))
	
	go UDP_receive(port, listen_ch, 1000)
	
	for i := range active_elevator_list {
		active_elevator_list[i] = IP_list[i]
	}
	
	for {
		message				:= <- listen_ch
		if message.ID == 1 {
			for i := 0; i < len(IP_list); i++ {
			        if message.Trunc_IP == IP_list[i] {
					auxilary_list[i] = 0
				} else {
					if active_elevator_list[i] != 0 {
						auxilary_list[i] = auxilary_list[i] + 1
					}
				}
				if auxilary_list[i] > 5 {
					fmt.Println("Elevator", IP_list[i], "is d-e-a-d DEAD")
					active_elevator_list[i] = 0
					auxilary_list[i] = 0
				}			
			}
		}
		time.Sleep(1000*time.Millisecond)
		active_elevator_list_ch <- active_elevator_list
	}
}

func Utilities_send_new_order(port string, new_order string, receive_ch chan Message) {
	
		fmt.Println("---Sending NEW ORDER -", new_order)
		var order Message
        var sleep_time int 		= 10
        
        terminate_ch 			:= make(chan int, 10)
        
        order.ID                	= 3 // send new order
        order.Order_type        	= new_order
        
        go UDP_broadcast("129.241.187.255:" + port, order, sleep_time, terminate_ch)
        
        for {
        	select {
        	case i := <- receive_ch:
        		if i.ID == ORDER_NEW_ACK {
        			fmt.Println("---Order", new_order, "has arived at Master")
        			
        			terminate_ch <- 1
        			return
        		}
        	}
        }
}

func Utilities_execute_order(port string, order string) {
	
}

func Utilities_ack_order(port string, message_ID int) {
	
	var message Message
	var sleep_time int 		= 10
	terminate_ch			:= make(chan int, 10)
	
	message.ID 			= message_ID

	go UDP_broadcast("129.241.187.255:" + port, message, sleep_time, terminate_ch)
	
	time.Sleep(1000*time.Millisecond)
	terminate_ch <- 1
}

func Utilities_send_order_done(message Message, port string, receive_ch chan Message) {
	
		fmt.Println("---Order", message.Order_type, "is done. Please remove it from the queue")
		var order Message
        var sleep_time int 			= 10
        
        terminate_ch 				:= make(chan int, 10)
        
        order.ID                	= ORDER_DONE // send new order
        order.Order_type        	= message.Order_type
        
        go UDP_broadcast("129.241.187.255:" + port, order, sleep_time, terminate_ch)
        
        for {
        	select {
        	case i := <- receive_ch:
        		if i.ID == ORDER_DONE_ACK {
        			fmt.Println("---Order transaction terminated")
        			terminate_ch <- 1
        			return
        		}
        	}
        }
}

func Utilities_listen_state_and_update_state_matrix(state_port string, IP_list []int, state_matrix [][]int) {
	
	var sleep_time int 		= 100
        receive_ch			:= make(chan Message, 10)
        
        go UDP_receive(state_port, receive_ch, sleep_time)
        
        for {
        	select {
        	case msg := <- receive_ch:
        		if msg.ID != 2 {
        			fmt.Println("Not a state on STATE channel!")	
        		}
        		for i := 0; i < len(IP_list); i++ {
        			if msg.Trunc_IP == IP_list[i] {
        				state_matrix[0][i] = msg.Direction
        				state_matrix[1][i] = msg.Latest_floor
        			}
        		}        		
        	}        	        	
        }
}











