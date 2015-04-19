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

type Message struct {

    ID 				int // 0 for init, 1 for <I am alive>, 9000 for error, 666 to escape current go routine
    Latest_floor 	int // 1 for first, 2 for second and so on
    Direction 		int	// 0 for NO_DIRECTION, 1 for UP and 2 for DOWN
    Current_time 	time.Time
    Local_IP 		string
    Remote_IP 		string 
}

func UDP_receive(port string, receive_ch chan Message, terminate_ch chan int)(int, error) {
	
	var receive_message Message
	
	fmt.Println("Creating new UDP_receive")
	baddr, err 			:= net.ResolveUDPAddr("udp",":"+port)
	if err != nil {fmt.Println(err); return -1, err}

	localListenConn, err := net.ListenUDP("udp", baddr)
	if err != nil {fmt.Println(err); return -1, err}
	
	for {
		select {
		case <- terminate_ch:
			fmt.Println("Terminating UDP_receive")
			localListenConn.Close()
			return 0, err
		default:
			buffer 			:= make([]byte, 1024)
		
			n,_,err 		:= localListenConn.ReadFromUDP(buffer[0:])
			if err != nil {fmt.Println(err); return -1, err}
		
			err 			= json.Unmarshal(buffer[:n], &receive_message)
			if err != nil {fmt.Println(err); return -1, err}
			receive_ch <- receive_message
			time.Sleep(500*time.Millisecond)
		}
	}
}

func UDP_broadcast(baddr string, msg Message, terminate_ch chan int, IP_ch chan string) (int, error){

	tempConn, err 		:= net.Dial("udp", baddr)
	if err != nil {fmt.Println(err); return -1, err}
	
	msg.Local_IP 		= tempConn.LocalAddr().String()
	msg.Remote_IP 		= baddr
	
	IP_ch <- msg.Local_IP
	
	for {
		select {
		case <- terminate_ch:
			fmt.Println("Terminating UDP_broadcast")
			return -1, err
		default:
			msg.Current_time 		= time.Now()
			buff, err 			:= json.Marshal(msg)
			if err != nil {fmt.Println(err); return -1, err}
			tempConn.Write([]byte(buff))
			time.Sleep(1000*time.Millisecond)
			fmt.Println("BROADCASTING :)")
		}
	}
	
	return 0, err
}

func Network_init(N int, port string)  ([]int, int){

	var msg Message
	init_ch 			:= make(chan Message, 1024)
	terminate_ch		:= make(chan int, 10)
	IP_ch				:= make(chan string, 100)
	
	msg.ID 				= 0													// init ID = 0
	
	go UDP_broadcast("129.241.187.255:" + port, msg, terminate_ch, IP_ch)
	go UDP_receive(port, init_ch, terminate_ch)
	
	IP_list 			:= Network_capture_IP(N, init_ch)	
	Temp_IP 			:= <- IP_ch
	terminate_ch 		<- 1
	terminate_ch 		<- 1
	time.Sleep(100*time.Millisecond)
	
	Utilities_bubble_sort(IP_list)
	Local_IP, _ := strconv.Atoi(Temp_IP[12:15])
	return IP_list, Local_IP
}

func Network_capture_IP(N int, init_ch chan Message) ([]int){
	
	IP_list 			:= make([]int, N)
	loop_counter 		:= 1
	
	
	
	for loop_counter  < N+1 {
	
		chan_value 			:= <- init_ch
		Local_IP, _ := strconv.Atoi(chan_value.Local_IP[12:15])
		i 				:= 0
		for j := 0; j < N; j++ {
			
			if Local_IP == IP_list[j] {
				i = 1
			}
		}
		if i != 1 && chan_value.ID == 0 {
			fmt.Println("\n________________________________________________________________________________")
			fmt.Println("NEW IP address!")
			fmt.Println("Pushing elevator nr.",loop_counter, "with ID", Local_IP, "into the list!")
			fmt.Println(loop_counter, "out of", N, "\n")

			IP_list[loop_counter-1] = Local_IP
			loop_counter++
		}
		time.Sleep(100*time.Millisecond)
	}
		
	fmt.Println("\n________________________________________________________________________________")
	fmt.Println("All(", N, ") IP addresses aquired\n")
	fmt.Println("Starting initialization of elevators\n")
	
	
	return IP_list
}




/* Replaced by two separate functions
func Network_init_2(N int, port string) {

	var msg Message
	init_channel 	:= make(chan Message, 1024)
	terminate_ch	:= make(chan int, 10)
	
		
	IP_list 		:= list.New() 
	IP_list.PushFront(N+1)
	
	msg.ID 			= 0													// init ID = 0
	msg.Current_time 	= time.Now()
	
	go UDP_broadcast("129.241.187.255:" + port, msg, terminate_ch)
	go UDP_receive(port, init_channel, terminate_ch)
	
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
			
			//Local_IP,_ := strconv.Atoi(i.Local_IP[12:15])
			Local_ID := IP_list.Len()
			IP_list.PushFront(Local_ID)
		}
		
		time.Sleep(100*time.Millisecond)
	}
	
	//Utilities_bubble_sort(IP_list)
	terminate_ch <- 1
	terminate_ch <- 1
	
	fmt.Println("\n________________________________________________________________________________")
	fmt.Println("All(", N, ") IP addresses aquired\n")
	fmt.Println("Starting initialization of elevators\n")	
}
*/







