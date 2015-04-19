// go run init.go

package main

import ("./driver"
		//"fmt"
		)

const n_elevators = 1
var port string 	= "26816"

func main() {

	driver.Network_init(n_elevators, port)
	
	driver.Elev_init()
	
	
	
	driver.Slave_run(n_elevators, port)
	/*
	driver.Elev_master()
	*/
}


