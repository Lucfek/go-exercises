package main

import (
	"fmt"
	"math/rand"
	"time"
)

func ping(c chan<- string) {

	time.Sleep(time.Duration(rand.Intn(100)*10) * time.Millisecond)
	c <- "ping"

}
func pong(c chan<- string) {
	time.Sleep(time.Duration(rand.Intn(100)*10) * time.Millisecond)
	c <- "pong"

}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	channel := make(chan string)

	for {
		go ping(channel)
		fmt.Println(<-channel)
		go pong(channel)
		fmt.Println(<-channel)

	}
}
