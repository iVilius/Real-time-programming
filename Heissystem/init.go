// go run init.go

package main

import ("./driver"
	"fmt")

const n_elevators 	= 1
const m_floors 		= 4
var init_port  string	= "19177"
var alive_port string 	= "47143"
var order_port string	= "80085"
var state_port string 	= "57473"

const	( 	
	INIT 		= 0
	ALIVE 		= 1
	STATE		= 2
	ORDER_NEW 	= 3
	ORDER_NEW_ACK 	= 4
	ORDER_ASSIGN	= 5
	ORDER_ASSIGN_ACK= 6
	)

func main() {
	
	if m_floors == 1 {
		fmt.Println("main: you don't need elevator for 1 floor!")
		return 	
	}	


	IP_list, Local_IP := driver.Network_init(n_elevators, init_port)

	driver.Elev_init(n_elevators, m_floors)
	
	Row := driver.Utilities_find_column_in_state_matrix(Local_IP, IP_list)
	
	if Local_IP == IP_list[0] {
		
		driver.Master_main(n_elevators, m_floors, alive_port, order_port, state_port, Row, IP_list)
	} else {
		driver.Slave_main(n_elevators, m_floors, alive_port, order_port, state_port, Row, IP_list)
	}
}


