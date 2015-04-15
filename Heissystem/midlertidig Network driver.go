// go run Exercise3_UDP.go

package main

import (
	"fmt"
	"net"
	"time"
)

func UDP_receive(port string, receive_ch chan string)(err error) {
	
	baddr, err := net.ResolveUDPAddr("udp",":"+port)
	
	if err != nil {
		return err
	}

	localListenConn, err := net.ListenUDP("udp", baddr)
	
	if err != nil {
		return err
	}

	for {
		buffer := make([]byte, 1024)
		localListenConn.ReadFromUDP(buffer[0:])
		receive_ch <- string(buffer)
	}
}

func UDP_send(baddr string, message string) (error){

	//var laddr *net.UDPAddr // local IP
	//baddr, err := net.ResolveUDPAddr("udp",":"+port)
	tempConn, err := net.Dial("udp", baddr)
	
	if err != nil {
		return err
	}
	
	for{
		tempConn.Write([]byte(message))
		time.Sleep(100*time.Millisecond)
	}
}

func main() {
	receive_ch := make(chan string, 100)

	go UDP_receive("20020", receive_ch)
	
	//go UDP_send("129.241.187.255:20020", "LFM UBRS")
	time.Sleep(100*time.Millisecond)

	for {
		i := <- receive_ch
		fmt.Println(i)
	}
}

