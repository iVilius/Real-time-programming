package driver

import ("fmt"
		"time")

func Slave_run(n_elevators int, m_floors int, port string, row int) {

	time.Sleep(2000*time.Millisecond)
	fmt.Print("I am slave\n")
	Slave_i_am_alive(n_elevators, m_floors, port, row)
	
}

func Slave_i_am_alive(n_elevators int, m_floors int, port string, row_in_state_matrix int) {
	
	var msg Message
	msg.ID 					= 1 // for <I am alive>
	msg.Latest_floor 		= Elev_get_latest_floor()
	msg.Direction 			= Elev_get_direction()
	
	terminate_ch			:= make(chan int, 1)
	door_ch					:= make(chan int, 1)
	order_ch 				:= make(chan [][]int, 5)
	IP_ch					:= make(chan string, 1)
	//stop_ch				:= make(chan bool, 1)
	
	order_ch <- Orders_make_state_matrix(n_elevators, m_floors)
	
	
	go UDP_broadcast("129.241.187.255:" + port, msg, terminate_ch, IP_ch)
	
	go Lamps_on(m_floors, row_in_state_matrix, door_ch, order_ch, terminate_ch)
	
	for {
		time.Sleep(1000*time.Millisecond)
	}
	//go Buttons_check(door_ch, order_ch)
	
	
	
	//error check
	/*
	slave_channel 	:= make(chan Message, 100)
	
	go UDP_receive(port, slave_channel, quit)

	
	for {
		go Lamps_run()
		i := <- slave_channel
		fmt.Println(i.ID)
		fmt.Println(i.FLO)
		fmt.Println(i.DIR)
		fmt.Println(i.Local_IP)
		fmt.Println(i.Remote_IP)
		fmt.Println(i.Curr_time)
	}*/
}

