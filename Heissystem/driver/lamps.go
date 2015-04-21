package driver

import ("fmt"
	"time"
	)

const (
	

)

func Lamps_check(M int, row int, door_ch chan int, lamp_ch chan [][]int, terminate_ch chan int) {
	
	for {
		time.Sleep(50*time.Millisecond)
		select {
		case i := <- door_ch:
			go Lamps_door(i)
		case i := <- lamp_ch:
			go Lamps_order_buttons(i, M, row)
		/*case i := <- stop_ch:
			go Lamps_stop()*/
		case <- terminate_ch:
			return
		default:
			Lamps_stop_button()
			Lamps_latest_floor()
			
			time.Sleep(50*time.Millisecond)
		}
	}
}

func Lamps_stop_button() {

	if IO_read_analog(STOP) == 1 {
		IO_write_analog(LIGHT_STOP, 1)
	} else {
		IO_write_analog(LIGHT_STOP, 0)
	}
}

func Lamps_latest_floor() {

	switch Elev_get_latest_floor() {
	case 1:
		IO_clear_bit(LIGHT_FLOOR_IND1)
		IO_clear_bit(LIGHT_FLOOR_IND2)
	case 2:
		IO_clear_bit(LIGHT_FLOOR_IND1)
		IO_set_bit(LIGHT_FLOOR_IND2)
	case 3:
		IO_set_bit(LIGHT_FLOOR_IND1)
		IO_clear_bit(LIGHT_FLOOR_IND2)
	case 4:
		IO_set_bit(LIGHT_FLOOR_IND1)
		IO_set_bit(LIGHT_FLOOR_IND2)
	default:
	}
}

func Lamps_order_buttons(array [][]int, m_floors int, row int) {
	
	if m_floors == 2 {
		IO_write_analog(ORDER1, array[row][3])
		IO_write_analog(ORDER2, array[row][4])
		IO_write_analog(UP1, 	array[row][5])
		IO_write_analog(DOWN2, 	array[row][6])
	} else if m_floors == 3 {
		IO_write_analog(ORDER1, array[row][3])
		IO_write_analog(ORDER2, array[row][4])
		IO_write_analog(ORDER3, array[row][5])	
		IO_write_analog(UP1, 	array[row][6])
		IO_write_analog(UP2, 	array[row][7])
		IO_write_analog(DOWN2, 	array[row][8])
		IO_write_analog(DOWN3, 	array[row][9])
	} else if m_floors == 4 {
		IO_write_analog(ORDER1, array[row][3])
		IO_write_analog(ORDER2, array[row][4])
		IO_write_analog(ORDER3, array[row][5])
		IO_write_analog(ORDER4, array[row][6])
		IO_write_analog(UP1, 	array[row][7])
		IO_write_analog(UP2, 	array[row][8])
		IO_write_analog(UP3, 	array[row][9])
		IO_write_analog(DOWN2, 	array[row][10])
		IO_write_analog(DOWN3,	array[row][11])
		IO_write_analog(DOWN4, 	array[row][12])
	} else {
		fmt.Println("Lamps: the system is not suited for", m_floors, "floors")
		return
	}
}


func Lamps_door(value int) {
	
	IO_write_analog(LIGHT_DOOR_OPEN, value)
}






