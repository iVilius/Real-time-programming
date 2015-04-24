package driver

import (//"fmt"
		"time")
		
func Sensors_get_latest_floor() int {
	
	if (IO_read_bit(SENSOR_FLOOR1) == 1) {
		return 1
	} else if (IO_read_bit(SENSOR_FLOOR2) == 1) {
		return 2
	} else if (IO_read_bit(SENSOR_FLOOR3) == 1) {
		return 3
	} else if (IO_read_bit(SENSOR_FLOOR4) == 1) {
		return 4
	} else {
		return -1 // error
	}
}

// SKAL TIL SENSORS

func Sensors_get_obstruction() {

	IO_read_bit(OBSTRUCTION)
}

// SKAL TIL SENSORS

func Sensors_get_stop() {

	IO_read_bit(STOP)
}
		
func Sensors_main(state_matrix [][]int, row int, receive_ch chan Message) {

	for {
		time.Sleep(50*time.Millisecond)
		select {
		
		/*case i := <- order_ch:
			go Sensors_buttons(i, row, M_floors)*/
		default:		
			time.Sleep(50*time.Millisecond)
			Sensors_buttons(state_matrix, row, receive_ch)
		}
	}
}


func Sensors_buttons(matrix [][]int, row int, receive_ch chan Message) {
	
	var message Message
	
	if IO_read_analog(BUTTON_COMMAND1) == 1 {
		message.Order_type = "1ORDER"
		if Orders_update_state_matrix(message, matrix, row, 1) == 0 {
			Utilities_send_new_order(Order_port, message.Order_type, receive_ch)
		}
	}
	if IO_read_analog(BUTTON_COMMAND2) == 1 {
		message.Order_type = "2ORDER"
		if Orders_update_state_matrix(message, matrix, row, 1) == 0 {
			Utilities_send_new_order(Order_port, message.Order_type, receive_ch)
		}
	}
	if IO_read_analog(BUTTON_COMMAND3) == 1 {
		message.Order_type = "3ORDER"
		if Orders_update_state_matrix(message, matrix, row, 1) == 0 {
			Utilities_send_new_order(Order_port, message.Order_type, receive_ch)
		}
	}
	if IO_read_analog(BUTTON_COMMAND4) == 1 {
		message.Order_type = "4ORDER"
		if Orders_update_state_matrix(message, matrix, row, 1) == 0 {
			Utilities_send_new_order(Order_port, message.Order_type, receive_ch)
		}
	}
	
	/*
	Orders_update_state_matrix(matrix, row, M_floors, 2, "ORDER", IO_read_analog(BUTTON_COMMAND2))
	Orders_update_state_matrix(matrix, row, M_floors, 3, "ORDER", IO_read_analog(BUTTON_COMMAND3))
	Orders_update_state_matrix(matrix, row, M_floors, 4, "ORDER", IO_read_analog(BUTTON_COMMAND4))
	Orders_update_state_matrix(matrix, row, M_floors, 1, "UP", IO_read_analog(BUTTON_UP1))
	Orders_update_state_matrix(matrix, row, M_floors, 2, "UP", IO_read_analog(BUTTON_UP1))
	Orders_update_state_matrix(matrix, row, M_floors, 3, "UP", IO_read_analog(BUTTON_UP1))
	Orders_update_state_matrix(matrix, row, M_floors, 2, "DOWN", IO_read_analog(BUTTON_DOWN2))
	Orders_update_state_matrix(matrix, row, M_floors, 3, "DOWN", IO_read_analog(BUTTON_DOWN2))
	Orders_update_state_matrix(matrix, row, M_floors, 4, "DOWN", IO_read_analog(BUTTON_DOWN2))
	matrix_back <- matrix*/
}
