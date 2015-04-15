// go run Exercise4.go

package driver

import (
	"fmt"
	"net"
	"time"
	"encoding/json"
	"container/list"
)

type Message struct {
    ID int
    curr_time time.Time
    local_IP string
    remote_IP string
}

func UDP_receive(port string, receive_ch chan Message)(err error) {
	
	baddr, err := net.ResolveUDPAddr("udp",":"+port)
	if err != nil {return err}

	localListenConn, err := net.ListenUDP("udp", baddr)
	if err != nil {return err}
	var receive_message Message
	
	for {
		time.Sleep(100*time.Millisecond)
		buffer := make([]byte, 2048)
		n,_,err := localListenConn.ReadFromUDP(buffer[0:])
		if err != nil {fmt.Println(err); return err}
		//fmt.Println(string(buffer))
		err = json.Unmarshal(buffer[:n], &receive_message)
		if err != nil {
			fmt.Println(err)
			return err
		}

		receive_ch <- receive_message
	}
}

func UDP_broadcast(baddr string, message Message) (error){

	tempConn, err := net.Dial("udp", baddr)
	if err != nil {return err}
	
	message.local_IP = tempConn.LocalAddr().String()
	message.remote_IP = baddr
	fmt.Println(message.local_IP)
	fmt.Println(message.remote_IP)
	fmt.Println(message.curr_time)
	//buffer := make([]byte, 2048)
	
	buffer, err := json.Marshal(message)
	if err != nil {return err}
	
	for{
		fmt.Println(string(buffer))
		tempConn.Write([]byte(buffer))
		time.Sleep(100*time.Millisecond)
	}
}

func Network_init(N int, port string) {
	var msg Message
	init_channel := make(chan Message, 1024)
	
	
	IP_list := list.New() 

	msg.ID = -1
	msg.remote_IP = ""
	msg.curr_time = time.Now()
	
	go UDP_broadcast("129.241.187.255:" + port, msg)
	go UDP_receive(port, init_channel)
	fmt.Println("Line 71\n")
	
	
	N = 1
	for IP_list.Len() < N {
		i := <- init_channel
		for element := IP_list.Front(); element != nil; element = element.Next() {
			if i.local_IP != element.Value {
				IP_list.PushFront(i.local_IP)
			}
		}
	}
	
	/*fmt.Println(i.ID)
	fmt.Println(i.local_IP)
	fmt.Println(i.remote_IP)
	fmt.Println(i.curr_time)*/
	fmt.Println("____________________________________\n")
	fmt.Println("All IP addresses aquired\n")
	fmt.Println("Starting initialization of elevators\n")	
	fmt.Println("____________________________________\n")
}

/*func main() {
	receiveChannel := make(chan Message, 1024)
	sendChannel := make(chan string, 1024)
	//message := Message{}
	go UDP_broadcast("129.241.187.255:30000", )
	go UDP_receive("30000", receiveChannel)
	
	time.Sleep(500*time.Millisecond)

	for {
		sendChannel <-"NOt generic"
		i := <- receiveChannel
		fmt.Println("\n\nMessage received on: ", i.CurrTime)
		fmt.Println("\nMessage ID was: ", i.ID)
		fmt.Println("\nMessage contents: ", i.Word)
		fmt.Println("\nLocal IP was: ", i.LocalIP)
		fmt.Println("\nRemote IP was: ", i.RemoteIP)
		fmt.Println("\nRaw contents: ", i.RawWord)
		fmt.Println("__________________________\n")
		time.Sleep(100*time.Millisecond)
	}
}
*/
