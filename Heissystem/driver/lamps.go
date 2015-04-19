package driver


const (
	

)

func Lamps_run(chan door_ch int, chan order_ch [][]int) {
	
	if IO_read_analog(STOP) == 1 {
		IO_write_analog(LIGHT_STOP, 1)
	} else {
		IO_write_analog(LIGHT_STOP, 0)
	}
	

}
