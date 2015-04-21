// go run Exercise3_2.go

package main

import (
	"fmt"
	"net"
	"time"
)

func TCP_accept(receive_ch2 chan string)(error){
	tcpAddr,err:=net.ResolveTCPAddr("tcp","129.241.187.155:21005")
	if err != nil {
			fmt.Println(err)
			return err
	}
	conn, err := net.ListenTCP("tcp", tcpAddr);
	
	if err != nil {
			fmt.Println(err)
			return err
	}
	
	tcp_conn, err := conn.Accept()
	if err != nil {
			fmt.Println(err)
			return err
	}
	
	go TCP_send(tcp_conn)
	go TCP_receive(tcp_conn, receive_ch2)
	
	return tcp_conn.Close()
}

func TCP_send(conn net.Conn)(error){
	for {
		_, err := conn.Write([]byte("SOmething not generic"))
		if err != nil {
			fmt.Println(err)
			return err
		}
		time.Sleep(100*time.Millisecond)
	}
}

func TCP_receive(conn net.Conn, receive_ch2 chan string)(error){
	for {
		buffer1 := make([]byte, 1024)
		_, err1 := conn.Read(buffer1[0:])
		if err1 != nil {
			fmt.Println(err1)
			return err1
		}
		receive_ch2 <- string(buffer1)
	}
	
}

func TCP_connect(receive_ch chan string, message string)(error){
	
	conn, err := net.Dial("tcp", "129.241.187.136:33546")
	
	if err != nil {
		fmt.Println(err)
		return err
	}
	
		conn.Write([]byte(message))
		time.Sleep(100*time.Millisecond)
		
		buffer := make([]byte, 1024)
		conn.Read(buffer[0:])
		receive_ch <- string(buffer)
	
	return conn.Close()
}

func main() {
	receive_ch := make(chan string, 1024)
	receive_ch2 := make(chan string, 1024)
	go TCP_connect(receive_ch, "Connect to: 129.241.187.155:21005\x00")
	go TCP_accept(receive_ch2)

	for{
			i := <- receive_ch
			k := <- receive_ch2
			fmt.Println(i)
			fmt.Println(k)
	}
}
