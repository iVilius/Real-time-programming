package driver

import ("fmt"
	"time")

func Master_main(alive_port string, order_port string, state_port string, row int, IP_list []int) {
        
    time.Sleep(2000*time.Millisecond)
	fmt.Print("I am master\n")
	
	var active_elevator_list 	[]int
	
	terminate_ch				:= make(chan int, 1)
	active_elevator_list_ch 	:= make(chan []int, 100)
	receive_ch					:= make(chan Message, 100)
	
	state_matrix := Orders_make_state_matrix()
	
	go UDP_receive(order_port, receive_ch, 10)
    go Utilities_send_i_am_alive(alive_port)
    go Lamps_main(row, state_matrix, terminate_ch)
	go Sensors_main(state_matrix, row, receive_ch)
        
    go Utilities_whos_alive(alive_port, IP_list, active_elevator_list_ch)
    go Utilities_listen_state_and_update_state_matrix(state_port, IP_list, state_matrix) 
    
    for {
    	select {
        case i := <- receive_ch:
        	if i.ID == ORDER_NEW{
        		Utilities_ack_order(order_port, ORDER_NEW_ACK)
        		fmt.Println("---I acknowledge new order request")
        		
        		active_elevator_list := <- active_elevator_list_ch
        		Cost_main(i, order_port, state_matrix, active_elevator_list, receive_ch)
        	} else if i.ID == ORDER_DONE {
        		Utilities_ack_order(order_port, ORDER_DONE_ACK)
        		
        		active_elevator_list = <- active_elevator_list_ch
        		Cost_delete_order(order_port, i, state_matrix, active_elevator_list)
        		Cost_send_state_matrix(state_matrix, order_port)
        	}
        default:
        	time.Sleep(10*time.Millisecond)

        }
    }
        
        /*for {
		active_elevator_list := <- active_elevator_list_ch
		fmt.Println(active_elevator_list)
		time.Sleep(1000*time.Millisecond)
	}  */      
}
