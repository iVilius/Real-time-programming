package driver

import ("fmt"
	"time")

func Slave_main(n_elevators int, m_floors int, port string, row int) {

	time.Sleep(2000*time.Millisecond)
	fmt.Print("I am slave\n")

	door_ch			:= make(chan int, 1)
	order_ch 		:= make(chan [][]int, 5)
	terminate_ch		:= make(chan int, 1)

	order_ch <- Orders_make_state_matrix(n_elevators, m_floors)
	
	go Slave_i_am_alive(n_elevators, m_floors, port)
	go Lamps_on(m_floors, row, door_ch, order_ch, terminate_ch)
	
	for {
		time.Sleep(1000*time.Millisecond)
	}
}

func Slave_i_am_alive(n_elevators int, m_floors int, port string) {
	
	var msg Message
	msg.ID 			= 1 // for <I am alive>
	msg.Latest_floor 	= Elev_get_latest_floor()
	msg.Direction 		= Elev_get_direction()
	terminate_ch		:= make(chan int, 1)

	go UDP_broadcast("129.241.187.255:" + port, msg, terminate_ch)
}

