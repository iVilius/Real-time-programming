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

    ID int
    Curr_time time.Time
    Local_IP string
    Remote_IP string 
}

func UDP_receive(port string, receive_ch chan Message)(err error) {
	
	var receive_message Message
	
	baddr, err 			:= net.ResolveUDPAddr("udp",":"+port)
	if err != nil {return err}

	localListenConn, err := net.ListenUDP("udp", baddr)
	if err != nil {return err}
	
	for {
		
		buffer 			:= make([]byte, 2048)
		
		n,_,err 		:= localListenConn.ReadFromUDP(buffer[0:])
		if err != nil {fmt.Println(err); return err}
		
		err 			= json.Unmarshal(buffer[:n], &receive_message)
		if err != nil {fmt.Println(err); return err}
		
		receive_ch <- receive_message
	}
}

func UDP_broadcast(baddr string, msg Message) (error){

	tempConn, err 		:= net.Dial("udp", baddr)
	if err != nil {return err}
	
	msg.Local_IP 		= tempConn.LocalAddr().String()
	msg.Remote_IP 		= baddr
	msg.Curr_time 		= time.Now()
	
	buff, err 			:= json.Marshal(msg)
	if err != nil {return err}
	
	for{
		tempConn.Write([]byte(buff))
		time.Sleep(100*time.Millisecond)
	}
}

func Network_init(N int, port string) {
	var msg Message
	init_channel 	:= make(chan Message, 1024)
	
	IP_list 		:= list.New() 
	IP_list.PushFront(-1)
	
	msg.ID 			= 0													// init ID = 0
	msg.Curr_time 	= time.Now()
	
	go UDP_broadcast("129.241.187.255:" + port, msg)
	go UDP_receive(port, init_channel)
	time.Sleep(500*time.Millisecond)
	
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
			fmt.Println("Pushing", i.Local_IP, "into the list!")
			fmt.Println(IP_list.Len(), "out of", N, "\n")
			
			Local_IP,_ := strconv.Atoi(i.Local_IP[12:15])
			IP_list.PushFront(Local_IP)
		}
		
		time.Sleep(100*time.Millisecond)
	}
	fmt.Printf("%d", IP_list.Front())
	fmt.Println("\n________________________________________________________________________________")
	fmt.Println("All(", N, ") IP addresses aquired\n")
	fmt.Println("Starting initialization of elevators\n")	
}
