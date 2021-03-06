---Go channels and Go Routine
package main

import (
	"fmt"
	"time"
)

func try(done chan bool) {
	fmt.Println("Channel done started")
	time.Sleep(4 * time.Second)
	fmt.Println("Channel awaked after sleep")
	done <- true
}
func main() {
	done := make(chan bool)
	fmt.Println("Main goroutine")
	go try(done)
	<-done
	fmt.Println("Data Recieved")
}
