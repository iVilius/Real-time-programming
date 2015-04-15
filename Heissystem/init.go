// go run init.go

package main

import ("./driver"
		)

var n_elevators int = 5
var port string = "26816"

func main() {

	driver.Network_init(n_elevators, port)
	
	//elev_init()
	
}


