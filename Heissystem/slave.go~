package driver

func Slave_run(n_elevators int, port string) {

	Slave_i_am_alive(port)
	for {}
}

func Slave_i_am_alive(port string) {
	
	var msg Message
	msg.ID = 1 // for <I am alive>
	msg.FLO = Elev_get_latest_floor()
	msg.DIR = Elev_get_direction()
	
	go UDP_broadcast("129.241.187.255:" + port, msg)
}

