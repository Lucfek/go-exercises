package main

import (
	"fmt"
	"math/rand"
	"time"
)

func ping(c chan string) {
	for {
		select {
		case msg := <-c:
			fmt.Println(msg)

		default:
			c <- "ping"
			time.Sleep(time.Duration(rand.Intn(100)*10) * time.Millisecond)
		}
	}
}
func pong(c chan string) {
	for {
		select {
		case msg := <-c:
			fmt.Println(msg)
		default:
			c <- "pong"
			time.Sleep(time.Duration(rand.Intn(100)*10) * time.Millisecond)
		}
	}
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	c := make(chan string)

	go ping(c)
	go pong(c)

	for {

	}

}
