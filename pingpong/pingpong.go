package main

import (
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
		print(<-channel, "\n")
		go pong(channel)
		print(<-channel, "\n")

	}
}
