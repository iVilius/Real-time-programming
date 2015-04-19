// go run Exercise4.go

package driver

import (
	"fmt"
	"net"
	"time"
	"encoding/json"
	"container/list"
	"strconv"
)

type Message struct {

    ID int // 0 for init, 1 for <I am alive>, 9000 for error, 666 to escape current go routine
    FLO int // 1 for first, 2 for second and so on
    DIR int	// 0 for NO_DIRECTION, 1 for UP and 2 for DOWN
    Curr_time time.Time
    Local_IP string
    Remote_IP string 
}

func UDP_receive(port string, receive_ch chan Message, quit_ch chan int)(int, error) {
	
	var receive_message Message
	
	fmt.Println("Creating new UDP_receive")
	baddr, err 			:= net.ResolveUDPAddr("udp",":"+port)
	if err != nil {fmt.Println(err); return -1, err}

	localListenConn, err := net.ListenUDP("udp", baddr)
	if err != nil {fmt.Println(err); return -1, err}
	
	for {
		select {
		case <- quit_ch:
			fmt.Println("Quitting UDP_receive")
			localListenConn.Close()
			return 0, err
		default:
			buffer 			:= make([]byte, 2048)
		
			n,_,err 		:= localListenConn.ReadFromUDP(buffer[0:])
			if err != nil {fmt.Println(err); return -1, err}
		
			err 			= json.Unmarshal(buffer[:n], &receive_message)
			if err != nil {fmt.Println(err); return -1, err}
			receive_ch <- receive_message
			time.Sleep(500*time.Millisecond)
		}
	}
}

func UDP_broadcast(baddr string, msg Message, quit_ch chan int) (int, error){

	tempConn, err 		:= net.Dial("udp", baddr)
	if err != nil {fmt.Println(err); return -1, err}
	
	msg.Local_IP 		= tempConn.LocalAddr().String()
	msg.Remote_IP 		= baddr
	
	
	
	
	for {
		select {
		case <- quit_ch:
			fmt.Println("Quitting UDP_broadcast")
			return -1, err
		default:
			msg.Curr_time 		= time.Now()
			buff, err 			:= json.Marshal(msg)
			if err != nil {fmt.Println(err); return -1, err}
			tempConn.Write([]byte(buff))
			time.Sleep(1000*time.Millisecond)
		}
	}
	
}

func Network_init(N int, port string) {

	var msg Message
	init_channel 	:= make(chan Message, 1024)
	quit_ch	:= make(chan int, 10)
	
		
	IP_list 		:= list.New() 
	IP_list.PushFront(-1)
	
	msg.ID 			= 0													// init ID = 0
	msg.Curr_time 	= time.Now()
	
	go UDP_broadcast("129.241.187.255:" + port, msg, quit_ch)
	go UDP_receive(port, init_channel, quit_ch)
	
	for IP_list.Len() < N+1 {	//N+1
	
		counter 	:= 0
		i 			:= <- init_channel
		
		for element := IP_list.Front(); element != nil; element = element.Next() {
			Local_IP,_ := strconv.Atoi(i.Local_IP[12:15])
			if Local_IP == element.Value {
				counter = 1
			}
		}
		if counter != 1 && i.ID == 0 {
			fmt.Println("\n________________________________________________________________________________")
			fmt.Println("NEW IP address!")
			fmt.Println("Pushing elevator nr.",IP_list.Len(), "with ID", i.Local_IP[12:15], "into the list!")
			fmt.Println(IP_list.Len(), "out of", N, "\n")
			
			Local_IP,_ := strconv.Atoi(i.Local_IP[12:15])
			IP_list.PushFront(Local_IP)
		}
		
		time.Sleep(100*time.Millisecond)
	}
	quit_ch <- 1
	quit_ch <- 1
	
	fmt.Println("\n________________________________________________________________________________")
	fmt.Println("All(", N, ") IP addresses aquired\n")
	fmt.Println("Starting initialization of elevators\n")	
}








