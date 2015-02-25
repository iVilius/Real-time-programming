package driver  // where "driver" is the folder that contains io.go, io.c, io.h, channels.go, channels.c and driver.go
/*
#cgo LDFLAGS: -lcomedi -lm
#include "io.h"
*/
import "C"
import "fmt"

func elev_init() {
	
	if (!io_init()) {
		fmt.Println("The hardware failed to initialize\n")
		return 1 // error
	
	elev_reset_lamps()
	elev_reset_orders()
	elev_set_floor(0)
		
	return 0 // success
}

func elev_set_direction(dir int) { // spør studass om navning, trenger vi speed?
									// husk å definere 0, 1, -1 som enum
	if (dir == 0) {
		io_write_analog(MOTOR, 0)
		return 0
		}
	else if (dir == 1) {
		io_clear_bit(MOTORDIR)
		io_write_analog(MOTOR, 2800)
		return 0
		}
	else if (dir == -1) {
		io_set_bit(MOTORDIR)
		io_write_analog(MOTOR, 2800)
		return 0
		}
	else {
		fmt.Println("The give direction is invalid\n")
		return 1} // error
}

func elev_get_direction() {

	return io_read_bit(MOTORDIR)
}

func elev_set_destination_floor(floor int) {
	
	if (elev_get_latest_floor() == floor) { //dependencies?!?
		elev_set_direction(0)
		return 0
		}
	
	if (floor < elev_get_latest_floor()) {
		elev_set_direction(-1)
		// vente på signal
		
	}
	else if (floor > elev_get_latest_floor()) {
		elev_set_direction(1)
		// vente på signal
	
	else {return 1} //error
}

func elev_get_latest_floor() {
	
	if (io_read_bit(SENSOR_FLOOR1))
		return 0
	else if (io_read_bit(SENSOR_FLOOR2))
		return 1
	else if (io_read_bit(SENSOR_FLOOR3))
		return 2
	else if (io_read_bit(SENSOR_FLOOR4))
		return 3
	else
		return -1 // error
}

func elev_threads() {
	go elev_thread_lamps(matrix)
	go elev_thread_sensors(matrix)
	go elev_thread_buttons(matrix)
}
