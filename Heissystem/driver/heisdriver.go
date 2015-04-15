package driver  // where "driver" is the folder that contains io.go, io.c, io.h, channels.go, channels.c and driver.go

import "fmt"


const (
	NO_DIRECTION = 0
	UP 			 = 1
	DOWN 		 = -1		
)

func elev_init() int {
	
	if (!IO_init()) {
		fmt.Println("The hardware failed to initialize\n")
		return 1 // error
	}
	
	//elev_reset_lamps() // DEFINE!
	//elev_reset_orders()
	elev_set_destination_floor(0)
		
	return 0 // success
}

func elev_set_direction(dir int) int { // spør studass om navning, trenger vi speed?

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
	} else {
		fmt.Println("The given direction is invalid\n")
		return 1
	} // error
}

func elev_get_direction() int {

	return (IO_read_bit(MOTORDIR))
}

// TO DO
func elev_set_destination_floor(floor int) int {
	
	if (elev_get_latest_floor() == floor) { //dependencies?!?
		elev_set_direction(NO_DIRECTION)
		return 0
	}
	
	if (floor < elev_get_latest_floor()) {
		elev_set_direction(DOWN)
		// vente på signal
		
	} else if (floor > elev_get_latest_floor()) {
		elev_set_direction(UP)
		// vente på signal
	} else {
		fmt.Println("Destination floor is invalid\n")
		return 1
	} //error
	return 1
}

func elev_get_latest_floor() int {
	
	if (IO_read_bit(SENSOR_FLOOR1) == 1) {
		return 0
	} else if (IO_read_bit(SENSOR_FLOOR2) == 1) {
		return 1
	} else if (IO_read_bit(SENSOR_FLOOR3) == 1) {
		return 2
	} else if (IO_read_bit(SENSOR_FLOOR4) == 1) {
		return 3
	} else {
		return -1 // error
	}
}

func elev_set_door_light(value int) {

	if (value == 1) {
		IO_set_bit(LIGHT_DOOR_OPEN)
	} else { 
		IO_clear_bit(LIGHT_DOOR_OPEN)
	}
}

func elev_get_obstruction() {

	IO_read_bit(OBSTRUCTION)
}

func elev_set_stop_light(value int) {

	if (value == 1) {
		IO_set_bit(LIGHT_STOP)
	} else {
		IO_clear_bit(LIGHT_STOP)
	}	
}

func elev_get_stop() {

	IO_read_bit(STOP)
}

func elev_set_UPDOWN_light(button int, value int) {

	if (value == 1) {
		IO_set_bit(button)
	} else {
		IO_clear_bit(button)
	}
}

func elev_set_order_light(button int, value int) {

	if (value == 1) {
		IO_set_bit(button)
	} else {
		IO_clear_bit(button)
	}
}

func elev_threads() {
	/*go elev_thread_motor(matrix)
	go elev_thread_sensors(matrix)
	go elev_thread_buttons(matrix)
	*/
}