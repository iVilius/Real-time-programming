// go run init.go

package main

import ("./driver"
		"fmt")

func main() {
	
	if driver.M_floors == 1 {
		fmt.Println("main: you don't need elevator for 1 floor!")
		return 	
	}	
	
	IP_list, Local_IP := driver.Network_init(driver.Init_port)

	driver.Motor_init()
	
	Row := driver.Utilities_find_column_in_state_matrix(Local_IP, IP_list)
	
	if Local_IP == IP_list[0] {
		
		driver.Master_main(driver.Alive_port, driver.Order_port, driver.State_port, Row, IP_list)
	} else {
		driver.Slave_main(driver.Alive_port, driver.Order_port, driver.State_port, Row, IP_list)
	}
}


