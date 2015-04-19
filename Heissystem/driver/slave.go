package driver

import (//"fmt"
		"time")

func Slave_run(n_elevators int, port string) {

	Slave_i_am_alive(n_elevators, port)

}

func Slave_i_am_alive(n_elevators int, port string) {
	
	var msg Message
	msg.ID 			= 1 // for <I am alive>
	msg.FLO 		= Elev_get_latest_floor()
	msg.DIR 		= Elev_get_direction()
	
	quit 			:= make(chan int, 10)
	
	time.Sleep(2000*time.Millisecond)
	
	go UDP_broadcast("129.241.187.255:" + port, msg, quit)
	
	go Orders_make_matrix(n_elevators)
	
	go Lamps_run(door_ch, order_ch)
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

