// go run init.go

package main

import ("./driver"
		//"fmt"
		)

const n_elevators 	= 1
const m_floors 		= 4
var port string 	= "26816"

func main() {
	
	IP_list, Local_IP := driver.Network_init(n_elevators, port)

	driver.Elev_init(n_elevators, m_floors)
	
	Row := driver.Utilities_find_collumn_in_state_matrix(Local_IP, IP_list)
	
	if Local_IP == IP_list[0] {
		
		driver.Slave_run(n_elevators, m_floors, port, Row)
	} else {
		//driver.Master_run(n_elevators, m_floors, port, Row)
	}
}


