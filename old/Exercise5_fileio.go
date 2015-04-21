// go run Exercise4.go

package main

import (
	"fmt"
	"time"
	"os"
	"strconv"
	"strings"
	"sync"
)

type Message struct {
    ID string
    Word string
   	CurrTime time.Time
    LocalIP string
    RemoteIP string
    RawWord string
}


func write_to_file(filename string, i int){
	fmt.Println("I am OPTIMUS Primary")
	// use flag os.O_APPEND|os.O_RDWR to append
	
	//Truncate a new file
	//f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0660)
	//err = f.Close()
	
	//Write to the file
	//if (err != nil) {fmt.Println(err)}
	
	var mutex = &sync.Mutex{}
	for {
		time.Sleep(100*time.Millisecond)
		mutex.Lock()
		f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0660)
		if (err != nil) {fmt.Println(err)}
		
		i++
		data1 := strconv.Itoa(i)
		data2 := []byte(data1)
		_, err1 := f.Write(data2)
		if (err1 != nil){fmt.Println(err)}
		
		//Close the file
		err = f.Close()
		if (err != nil) {fmt.Println(err)}
		mutex.Unlock()
	}
}


func read_from_file(filename string, sendCh chan int) {

	var mutex = &sync.Mutex{}
	for {
		mutex.Lock()
		f, err := os.Open(filename)
		//f, err := os.OpenFile(filename, os.O_RDWR, 0660)
		if (err != nil) {fmt.Println(err)}
		
		
		data2 := make([]byte, 100)
		
		n, err1 := f.Read(data2)
		if (err1 != nil){fmt.Println("Error at reading\n")}
		
		data1, err2 := strconv.Atoi(strings.TrimSuffix(string(data2[:n]),"\x00"))
		
		if (err2 != nil){fmt.Println(err2)}
		sendCh <- data1	
		
		err = f.Close()
		if (err != nil) {fmt.Println(err)}
		mutex.Unlock()
	}

}


func main() {

	var compare int
	var current int
	var counter int
	var i int
	
	const filename string = "Counter.txt"
	
	compare = 0
	for {
		f, err := os.Open(filename)
		//f, err := os.OpenFile(filename, os.O_RDWR, 0660)
		if (err != nil) {fmt.Println(err)}
		
		
		data2 := make([]byte, 100)
		
		n, err1 := f.Read(data2)
		if (err1 != nil){fmt.Println("Error at reading\n")}
		
		data1, err2 := strconv.Atoi(strings.TrimSuffix(string(data2[:n]),"\x00"))
		
		if (err2 != nil){fmt.Println(err2)}
		current = data1	
		
		err = f.Close()
		if (err != nil) {fmt.Println(err)}

		if (current == compare) {
			counter++
		} else {
			counter = 0
			compare = current
		}
		if (counter == 5) {
			i = current
			fmt.Println("Sup. Im Optimus Primary")
			for {
				
				f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0660)
				if (err != nil) {fmt.Println(err)}
		
				i++
				fmt.Println(i)
				data1 := strconv.Itoa(i)
				data2 := []byte(data1)
				_, err1 := f.Write(data2)
				if (err1 != nil){fmt.Println(err)}
		
				//Close the file
				err = f.Close()
				if (err != nil) {fmt.Println(err)}
				time.Sleep(100*time.Millisecond)
			}
		}
		time.Sleep(500*time.Millisecond)
		fmt.Println("Sup. Im a waste of recources")
	}
}

