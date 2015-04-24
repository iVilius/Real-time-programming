package driver

import ("fmt"
		"time")

func Slave_main(Alive_port string, Order_port string, state_port string, row int, IP_list []int, receive_ch chan Message) {

	time.Sleep(2000*time.Millisecond)
	fmt.Print("I am slave\n")

	terminate_ch				:= make(chan int, 1)
	active_elevator_list_ch 	:= make(chan []int, 100)
	//receive_ch					:= make(chan Message, 500)
	
	state_matrix 				:= Orders_make_state_matrix()
	
	//go UDP_receive(Order_port, receive_ch, 100)
	go Utilities_send_i_am_alive(Alive_port)
	go Lamps_main(row, state_matrix, terminate_ch)
	go Sensors_main(state_matrix, row, receive_ch)
	
	go Utilities_whos_alive(Alive_port, IP_list, active_elevator_list_ch)
	go Slave_send_state_to_master(state_port)
	
	for {
		select {
		case i := <- receive_ch:
			if i.ID == ORDER_ASSIGN {
				fmt.Println("I will execute order", i.Order_type)
				Utilities_ack_order(Order_port, ORDER_ASSIGN_ACK)
				time.Sleep(5000*time.Millisecond)
				Utilities_send_order_done(i, Order_port, receive_ch)
			}	
		default:
			i := <- receive_ch
			fmt.Println("RECEIVE", i.ID)
		}
	}
	
}

func Slave_send_state_to_master(port string) {

	var latest_floor int 		= Sensors_get_latest_floor()
	var current_direction int 	= Motor_get_direction()

	for { //Sjekker om endring i tilstand
		if (Sensors_get_latest_floor() != latest_floor || Motor_get_direction() != current_direction) {
			Utilities_broadcast_state(port)
			fmt.Println("New floor or direction!")
			time.Sleep(100*time.Millisecond)
		}
		latest_floor 		= Sensors_get_latest_floor()
		current_direction 	= Motor_get_direction()
		time.Sleep(100*time.Millisecond)
	}
}
