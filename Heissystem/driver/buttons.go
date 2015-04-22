package driver


import (//"fmt"
		"time")
		
func Sensors_poll(matrix_ch chan [][]int, row int, m_floors int) {

	for {
		time.Sleep(50*time.Millisecond)
		select {
		case i := <- matrix_ch:
			//fmt.Println(i)
			Sensors_buttons(matrix_ch, i, row, m_floors)
		/*case i := <- order_ch:
			go Sensors_buttons(i, row, m_floors)*/
		default:		
			time.Sleep(50*time.Millisecond)
		}
	}
}


func Sensors_buttons(matrix_ch chan [][]int, matrix [][]int, row int, m_floors int) {
	//fmt.Println(matrix)
	if IO_read_analog(BUTTON_COMMAND1) == 1 {
		Orders_update_state_matrix(matrix_ch, matrix, row, m_floors, 1, "ORDER", 1)
	}
	if IO_read_analog(BUTTON_COMMAND2) == 1 {
		Orders_update_state_matrix(matrix_ch, matrix, row, m_floors, 2, "ORDER", 1)
	}
	if IO_read_analog(BUTTON_COMMAND3) == 1 {
		Orders_update_state_matrix(matrix_ch, matrix, row, m_floors, 3, "ORDER", 1)
	}
	if IO_read_analog(BUTTON_COMMAND4) == 1 {
		Orders_update_state_matrix(matrix_ch, matrix, row, m_floors, 4, "ORDER", 1)
	}
	
	/*
	Orders_update_state_matrix(matrix, row, m_floors, 2, "ORDER", IO_read_analog(BUTTON_COMMAND2))
	Orders_update_state_matrix(matrix, row, m_floors, 3, "ORDER", IO_read_analog(BUTTON_COMMAND3))
	Orders_update_state_matrix(matrix, row, m_floors, 4, "ORDER", IO_read_analog(BUTTON_COMMAND4))
	Orders_update_state_matrix(matrix, row, m_floors, 1, "UP", IO_read_analog(BUTTON_UP1))
	Orders_update_state_matrix(matrix, row, m_floors, 2, "UP", IO_read_analog(BUTTON_UP1))
	Orders_update_state_matrix(matrix, row, m_floors, 3, "UP", IO_read_analog(BUTTON_UP1))
	Orders_update_state_matrix(matrix, row, m_floors, 2, "DOWN", IO_read_analog(BUTTON_DOWN2))
	Orders_update_state_matrix(matrix, row, m_floors, 3, "DOWN", IO_read_analog(BUTTON_DOWN2))
	Orders_update_state_matrix(matrix, row, m_floors, 4, "DOWN", IO_read_analog(BUTTON_DOWN2))
	matrix_back <- matrix*/
}
