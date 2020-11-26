---Ping Pong

package main

import (
	"fmt"
	"time"
)

func ping(click chan string) {
	for i := 1; i < 6; i++ {
		click <- "ping"
	}
}

func sleep(click chan string) {
	for {
		msg := <-click
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func pong(click chan string) {
	for i := 1; i < 6; i++ {
		click <- "pong"
	}
}

func main() {
	var click chan string = make(chan string)

	go ping(click)
	go sleep(click)
	go pong(click)

	var input string
	fmt.Scanln(&input)
}
