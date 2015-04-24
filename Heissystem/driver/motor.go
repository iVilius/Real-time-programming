package driver  // where "driver" is the folder that contains io.go, io.c, io.h, channels.go, channels.c and driver.go

import "fmt"


const (
	NO_DIRECTION 	= 0
	UP 				= 1
	DOWN 			= 2		
)

func Motor_init() int {
	
	
	/*
	DO THIS IF TIME LEFT
	
	door_ch 		:= make(chan int, 1)
	order_ch 		:= make(chan [][]int, 5)
	terminate_ch 	:= make(chan int, 1)
	go Lamps_on(door_ch, order_ch, terminate_ch) // Lamps_reset
	*/
	
	if (!IO_init()) {
		fmt.Println("The hardware failed to initialize\n")
		return 1
	}
	
	Motor_set_direction(DOWN)
	
	for {
		if Sensors_get_latest_floor() == 1 {
			Motor_set_direction(NO_DIRECTION)
			break
		}
	}

	fmt.Println("\n________________________________________________________________________________")
	fmt.Println("Initialization successful\n")	
	return 0 
}

func Motor_main() {
	
	
	
}

func Motor_set_direction(dir int) int { // spør studass om navning, trenger vi speed?

	if dir == NO_DIRECTION {
		IO_write_analog(MOTOR, 0)
		return 0
	} else if dir == UP {
		IO_clear_bit(MOTORDIR)
		IO_write_analog(MOTOR, 2800)
		return 0
	} else if dir == DOWN {
		IO_set_bit(MOTORDIR)
		IO_write_analog(MOTOR, 2800)
		return 0
	}
	
	fmt.Println("The given direction is invalid\n")
	return -1
}

func Motor_get_direction() int {
	
	if IO_read_bit(MOTOR) == 0 {
		return 0
	} else if IO_read_bit(MOTORDIR) == 1 {
		return 2
	} else if IO_read_bit(MOTORDIR) == 0 {
		return 1
	} else {
		fmt.Println("The direction is invalid\n")
	}
	return -1
}

func Motor_set_destination_floor(floor int) int {
	
	if (Sensors_get_latest_floor() == floor) { //dependencies?!?
		Motor_set_direction(NO_DIRECTION)
		return 0
	}
	
	if (floor < Sensors_get_latest_floor()) {
		Motor_set_direction(DOWN)
		for Sensors_get_latest_floor() != floor {}
		Motor_set_direction(NO_DIRECTION)
		return 0
	} else if (floor > Sensors_get_latest_floor()) {
		Motor_set_direction(UP)
		for Sensors_get_latest_floor() != floor {}
		Motor_set_direction(NO_DIRECTION)
		return 0
		
	} else {
		fmt.Println("Destination floor is invalid\n")
		return 1
	} //error
	return 1
}


func Elev_threads() {
	/*go elev_thread_motor(matrix)
	go elev_thread_sensors(matrix)
	go elev_thread_buttons(matrix)
	*/
}

//Lage KØ
//Lage en intern prioritetsalgoritme lik den vi hadde i datastyring
// Sjekke køen: Noen ordre i current direction? etc etc etc
