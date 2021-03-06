// go run Exercise4.go

package main

import (
	"fmt"
	"net"
	"time"
	"encoding/json"
)

type Message struct {
    ID string
    Word string
   	CurrTime time.Time
    LocalIP string
    RemoteIP string
    RawWord string
}

func UDP_receive(port string, receiveCh chan int, intChannel chan int)(err error) {
	
	baddr, err := net.ResolveUDPAddr("udp",":"+port)
	if err != nil {return err}

	localListenConn, err := net.ListenUDP("udp", baddr)
	if err != nil {return err}
	var receiveMessage int
	for {
		buffer := make([]byte, 2048)
		n,_,err := localListenConn.ReadFromUDP(buffer[0:])
		if err != nil {fmt.Println(err); return err}
		//fmt.Println(string(buffer))
		err = json.Unmarshal(buffer[:n], &receiveMessage)
		//receiveMessage.LocalIP = addr.String()
		//receiveMessage.CurrTime = time.Now()
		if err != nil {
			fmt.Println(err)
			return err
		}
		//receiveMessage.RawWord = string(buffer)
		receiveCh <- receiveMessage
		fmt.Println(receiveMessage)
		interrupt := <- intChannel
		if (interrupt == 1){
			fmt.Println("TERMINATING LISTENING")
			return err
		}
	}
}

func UDP_broadcast(baddr string, sendCh chan int) (error){
	tempConn, err := net.Dial("udp", baddr)
	if err != nil {return err}
	
	var msg int
	//msg.ID = "1"
	msg = <- sendCh
	//msg.RemoteIP = baddr

	buffer, err := json.Marshal(msg)
	if err != nil {return err}
	
	
	for{
		tempConn.Write([]byte(buffer))
		
		time.Sleep(500*time.Millisecond)
	}
}

/*func process_pair(sendChannel chan int, receiveChannel chan int, compChannel chan int) {
	var flag int
	for{
		i := <- receiveChannel
		fmt.Println(i)
		if i < 0 {
			flag = 0
			break
		
		} else{ 
			flag = 1
			break
		}
	}
	
	if (flag == 0) {
		go primary(0, sendChannel)
	} else {
		go secondary(sendChannel, receiveChannel, compChannel)
	}
}*/

func primary(i int, sendChannel chan int, intChannel chan int) {
	fmt.Println("I am primary!\n")
	go UDP_broadcast("129.241.187.255:24568", sendChannel)
	for {
		intChannel <- 1
		time.Sleep(100*time.Millisecond)
		i++
		sendChannel <- i
		//fmt.Println(i)
	}
	intChannel <- 0
}

func secondary(i int, sendChannel chan int, receiveChannel chan int, compChannel chan int, intChannel chan int) {
	
	//check if primary is running
	
	
}



func main() {
	var i int
	receiveChannel := make(chan int, 1024)
	sendChannel := make(chan int, 1024)
	compChannel := make(chan int, 1024)
	intChannel := make(chan int, 1024)
	sendChannel <- 0
	intChannel <- 0
	interrupt := <- intChannel
	
	go secondary(i, sendChannel, receiveChannel, compChannel, intChannel)
	
	
	time.Sleep(100*time.Millisecond)

	
	for {
		<- receiveChannel
		/*fmt.Println("\n\nMessage received on: ", i.CurrTime)
		fmt.Println("\nMessage ID was: ", i.ID)
		fmt.Println("\nMessage contents: ", i.Word)
		fmt.Println("\nLocal IP was: ", i.LocalIP)
		fmt.Println("\nRemote IP was: ", i.RemoteIP)
		fmt.Println("\nRaw contents: ", i.RawWord)
		fmt.Println("__________________________\n")*/
		//time.Sleep(100*time.Millisecond)
	}
}

