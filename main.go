package main

import (
	"fmt"
	"time"
)


func calc(workerId string, msg chan int){
	for  res := range msg {
		fmt.Printf("worker %s recebeu %d\n", workerId, res)
		time.Sleep(time.Second)
	}
}

func main() {
	// channel create
	msg := make(chan int)

	// thread 1 
	for i := 0; i < 10; i ++ {
		go calc(fmt.Sprintf("Worker %d", i), msg);
	}

	// thread 2 
	for i := 0; i < 100; i ++ {
		msg <- i
	}
}