package driver

import ("fmt"
		"time")

func Slave_main(n_elevators int, m_floors int, alive_port string, order_port string, state_port string, row int, IP_list []int) {

	time.Sleep(2000*time.Millisecond)
	fmt.Print("I am slave\n")

	door_ch					:= make(chan int, 1)
	terminate_ch				:= make(chan int, 1)
	matrix_ch				:= make(chan [][]int, 5)
	active_elevator_list_ch 		:= make(chan []int, 100)
	
	matrix_ch <- Orders_make_state_matrix(n_elevators, m_floors)
	matrix_ch <- Orders_make_state_matrix(n_elevators, m_floors)
	
	go Utilities_send_i_am_alive(n_elevators, m_floors, alive_port)
	go Lamps_check(m_floors, row, door_ch, matrix_ch, terminate_ch)
	go Sensors_poll(matrix_ch, row, m_floors)
	
	Utilities_whos_alive(alive_port, IP_list, active_elevator_list_ch)
	Slave_send_state_to_master(n_elevators, m_floors, state_port)
	
	for {
		active_elevator_list := <- active_elevator_list_ch
		fmt.Println(active_elevator_list)
		time.Sleep(1000*time.Millisecond)
	}
	
}

func Slave_send_state_to_master(n_elevators int, m_floors int, port string) {

	var latest_floor int 		= Elev_get_latest_floor()
	var current_direction int 	= Elev_get_direction()

	for { //Sjekker om endring i tilstand
		if (Elev_get_latest_floor() != latest_floor || Elev_get_direction() != current_direction) {
			Utilities_broadcast_state(n_elevators, m_floors, port)
			fmt.Println("New floor or direction!")
			time.Sleep(100*time.Millisecond)
		}
		latest_floor 		= Elev_get_latest_floor()
		current_direction 	= Elev_get_direction()
		time.Sleep(100*time.Millisecond)
	}
}
