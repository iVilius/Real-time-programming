package driver


func Master_main(n_elevators int, m_floors int, port string, row int, IP_list []int) {
        
        go Utilities_i_am_alive(n_elevators, m_floors, port)
        go Lamps_check(m_floors, row, door_ch, lamp_ch, terminate_ch)
	go Sensors_poll(matrix_ch, row, m_floors)
        
        Utilities_listen(port, IP_list, active_elevator_list_ch)
        
        
}
