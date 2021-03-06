// go run Exercise4.go

package driver

import (
	"fmt"
	"net"
	"time"
	"encoding/json"
	//"container/list"
	"strconv"
)

func UDP_receive(port string, receive_ch chan Message, sleep_time int)(int, error) {
	
	var receive_message Message
	
	fmt.Println("Creating new UDP_receive")
	baddr, err 				:= net.ResolveUDPAddr("udp",":"+port)
	if err != nil {fmt.Println(err); return -1, err}

	localListenConn, err 	:= net.ListenUDP("udp", baddr)
	if err != nil {fmt.Println(err); return -1, err}
	
	buffer 					:= make([]byte, 1024)
	
	n,_,err 				:= localListenConn.ReadFromUDP(buffer[0:])
	if err != nil {fmt.Println(err); return -1, err}
	
	err 					= json.Unmarshal(buffer[:n], &receive_message)
	if err != nil {fmt.Println(err); return -1, err}

	for {
		switch receive_message.ID {
		case -1:
			fmt.Println("Terminating UDP_receive")
			localListenConn.Close()
			return -1, err
			
		case 2:
			buffer 			:= make([]byte, 1024)
		
			n,_,err 		:= localListenConn.ReadFromUDP(buffer[0:])
			if err != nil {fmt.Println(err); return -1, err}
		
			err 			= json.Unmarshal(buffer[:n], &receive_message)
			if err != nil {fmt.Println(err); return -1, err}
			receive_ch <- receive_message
			fmt.Println("Received message with ID", receive_message.ID)
			//fmt.Println("Terminating UDP_receive")
			//time.Sleep(time.Duration(sleep_time)*time.Millisecond) 
			//return 2, err
		default:
			buffer 			:= make([]byte, 1024)
		
			n,_,err 		:= localListenConn.ReadFromUDP(buffer[0:])
			if err != nil {fmt.Println(err); return -1, err}
		
			err 			= json.Unmarshal(buffer[:n], &receive_message)
			if err != nil {fmt.Println(err); return -1, err}
			
			receive_ch <- receive_message
			/*
			if receive_message.ID == 1 {
				fmt.Println("I am alive!") }
				*/
			time.Sleep(time.Duration(sleep_time)*time.Millisecond)			
		}
	}
}

func UDP_broadcast(baddr string, msg Message, sleep_time int, terminate_ch chan int) (int, error){

	tempConn, err 		:= net.Dial("udp", baddr)//"129.241.187.255:" + baddr)
	if err != nil {fmt.Println(err); return -1, err}
	
	msg.Local_IP 		= tempConn.LocalAddr().String()
	msg.Remote_IP 		= baddr
	
	start_index			:= 12
	end_index			:= len(msg.Local_IP) - 6
	Trunc_IP, _ 		:= strconv.Atoi(msg.Local_IP[start_index:end_index])
	msg.Trunc_IP 		= Trunc_IP
	
	for {
		select {
		case <- terminate_ch:
			if msg.ID == 0 {
				msg.ID				= -1
				buff, err 			:= json.Marshal(msg)
				if err != nil {fmt.Println(err); return -1, err}
				tempConn.Write([]byte(buff))
				time.Sleep(time.Duration(500)*time.Millisecond)
				fmt.Println("Terminating UDP_broadcast")
				return -1, err
			}
			return -1, err
		default:
		}	
		switch msg.ID {
		case 2:
			buff, err 			:= json.Marshal(msg)
			if err != nil {fmt.Println(err); return -1, err}
			tempConn.Write([]byte(buff))
			time.Sleep(time.Duration(sleep_time)*time.Millisecond)
			msg.ID = 2
			fmt.Println("Terminating UDP_broadcast")
			return 2, err
			
		default:
			buff, err 			:= json.Marshal(msg)
			if err != nil {fmt.Println(err); return -1, err}
			tempConn.Write([]byte(buff))
			time.Sleep(time.Duration(sleep_time)*time.Millisecond)
		}
	}
	return 0, err
}

func Network_init(port string)  ([]int, int){

	var msg Message
	var sleep_time int
	
	terminate_ch		:= make(chan int, 5)
	init_ch 			:= make(chan Message, 1024)
	
	msg.ID 				= 0
	sleep_time			= 10
	
	go UDP_broadcast("129.241.187.255:" + port, msg, sleep_time, terminate_ch)
	go UDP_receive(port, init_ch, sleep_time)
	
	IP_list, Local_IP 	:= Network_capture_IP(init_ch)
	
	time.Sleep(1000*time.Millisecond)
	terminate_ch		<- -1	// Kill ongoing UDP processes
	

		
	Utilities_bubble_sort_asc(IP_list)

	return IP_list, Local_IP
}

func Network_capture_IP(init_ch chan Message) ([]int, int){
	
	IP_list 		:= make([]int, N_elevators)
	loop_counter 	:= 1
	
	for loop_counter  < N_elevators+1 {
				
		chan_value := <- init_ch
		i 		:= 0
		for j := 0; j < N_elevators; j++ {
			if chan_value.Trunc_IP == IP_list[j] {
					i = 1
				}
		}

		if i != 1 && chan_value.ID == 0 {
			fmt.Println		("\n________________________________________________________________________________")
			fmt.Println("NEW IP address!")
			fmt.Println("Pushing elevator nr.",loop_counter, "with ID", chan_value.Trunc_IP, "into the list!")
			fmt.Println(loop_counter, "out of", N_elevators, "\n")

			IP_list[loop_counter-1] = chan_value.Trunc_IP
			loop_counter++
		}
		time.Sleep(10*time.Millisecond)
	}
		
	fmt.Println("\n________________________________________________________________________________")
	fmt.Println("All(", N_elevators, ") IP addresses aquired\n")
	fmt.Println("Starting initialization of elevators\n")
	chan_value 	:= <- init_ch
	//start_index	:= 12
	//end_index	:= len(chan_value.Local_IP) - 6
	//Local_IP, _ 	:= strconv.Atoi(chan_value.Local_IP[start_index:end_index])	

	
	return IP_list, chan_value.Trunc_IP
}
