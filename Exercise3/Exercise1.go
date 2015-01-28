package main

import (
	. "fmt" // Using '.' to avoid prefixing functions with their 		package names
	// This is probably not a good idea for large projects...
	"runtime"
)

var i int


func thread_1(someChannel1 chan int, someChannel2 chan string) {
	for j := 0; j < 1000000; j++{
		<- someChannel1
		i++	
		someChannel1 <- 1
	}
	someChannel2 <- "Hey, I am done!"
}

func thread_2(someChannel1 chan int, someChannel2 chan string) {

	for k := 0; k < 1000001; k++{
		<- someChannel1
		i--
		someChannel1 <- 1
	}
	someChannel2 <- "Hey, I am done!"
}

func main(){
	i = 0
	//var key int

	runtime.GOMAXPROCS(runtime.NumCPU())

	someChannel1 := make(chan int, 1)
	someChannel1 <- 1
	someChannel2 := make(chan string, 2)

	go thread_1(someChannel1, someChannel2)
	go thread_2(someChannel1, someChannel2)
	
	<- someChannel2
	<- someChannel2

	Println(i)
}
