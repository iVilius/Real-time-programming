// go run Exercise3_1.go

package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

var laddr *net.UDPAddr //Local address
var baddr *net.UDPAddr //Broadcast address

func udp_init(receive_ch chan string) (err error) {
	//Generating broadcast address
	baddr, err = net.ResolveUDPAddr("udp4", "255.255.255.255:"+strconv.Itoa(30000))
	if err != nil {
		return err
	}

	//Generating localaddress
	tempConn, err := net.DialUDP("udp4", nil, baddr)
	defer tempConn.Close()
	tempAddr := tempConn.LocalAddr()
	laddr, err = net.ResolveUDPAddr("udp4", tempAddr.String())
	
	//Creating local listening connections
	localListenConn, err := net.ListenUDP("udp4", laddr)
	if err != nil {
		return err
	}

	//Creating listener on broadcast connection
	broadcastListenConn, err := net.ListenUDP("udp", baddr)
	if err != nil {
		localListenConn.Close()
	return err
	}

	go udp_receive_server(broadcastListenConn, receive_ch)
	
	return err
}

func udp_connection_reader(conn *net.UDPConn, rcv_ch chan string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("ERROR in udp_connection_reader: %s \n Closing connection.", r)
			conn.Close()
		}
	}()

	buf := make([]byte, 10)

	for {
		fmt.Printf("udp_connection_reader: Waiting on data from UDPConn\n")
		n, raddr, err := conn.ReadFromUDP(buf)
		fmt.Printf("udp_connection_reader: Received %s from %s \n", string(buf), raddr.String())
		if err != nil || n < 0 {
			fmt.Printf("Error: udp_connection_reader: reading\n")
			panic(err)
		}
		rcv_ch <- string(buf) // string(buf)
	}
}

func udp_receive_server(bconn *net.UDPConn, receive_ch chan string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("ERROR in udp_receive_server: %s \n Closing connection.", r)
			bconn.Close()
		}
	}()

	bconn_rcv_ch := make(chan string)

	go udp_connection_reader(bconn, bconn_rcv_ch)
	
	buf := <-bconn_rcv_ch
	receive_ch <- buf
}

func main() {

	
	receive_ch := make(chan string, 10)
	var bconn *net.UDPConn		

	go udp_init(receive_ch)
	go udp_receive_server(bconn, receive_ch)
	time.Sleep(100*time.Millisecond)
	
	i := <- receive_ch
	fmt.Printf(i)
}
