package driver

import ("fmt"
	"time")

func Master_main(n_elevators int, m_floors int, alive_port string, order_port string, state_port string, row int, IP_list []int) {
        
        time.Sleep(2000*time.Millisecond)
	fmt.Print("I am master\n")
	
	door_ch					:= make(chan int, 1)
	terminate_ch				:= make(chan int, 1)
	matrix_ch				:= make(chan [][]int, 5)
	active_elevator_list_ch 		:= make(chan []int, 100)
	state_matrix_ch				:= make(chan [][]int, 500)
	receive_ch				:= make(chan Message, 100)
	
	state_matrix := Orders_make_state_matrix(n_elevators, m_floors)
	matrix_ch <- Orders_make_state_matrix(n_elevators, m_floors)
	matrix_ch <- Orders_make_state_matrix(n_elevators, m_floors)
	
        go Utilities_send_i_am_alive(n_elevators, m_floors, alive_port)
        go Lamps_check(m_floors, row, door_ch, matrix_ch, terminate_ch)
	go Sensors_poll(matrix_ch, row, m_floors)
        
        Utilities_whos_alive(alive_port, IP_list, active_elevator_list_ch)
        Utilities_listen_state_and_update_state_matrix(state_port, IP_list, state_matrix, state_matrix_ch) 
        
        go UDP_receive(order_port, receive_ch, 100)
        
        for {
        	select {
        	case i := <-receive_ch:
        		if i.ID == 3 {
        			active_elevator_list := <- active_elevator_list_ch
        			Cost_main(i, order_port, state_matrix, active_elevator_list)
        		}
        	default:
        		time.Sleep(1000*time.Millisecond)
        	}
        }
        
        /*for {
		active_elevator_list := <- active_elevator_list_ch
		fmt.Println(active_elevator_list)
		time.Sleep(1000*time.Millisecond)
	}  */      
}
