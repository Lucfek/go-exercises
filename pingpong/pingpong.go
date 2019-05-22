package main

import (
	"math/rand"
	"time"
)

func ping(c chan<- string, waitTime int) {
	time.Sleep(time.Duration(waitTime) * time.Millisecond)
	c <- "ping"

}
func pong(c chan<- string, waitTime int) {
	time.Sleep(time.Duration(waitTime) * time.Millisecond)
	c <- "pong"

}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	channel := make(chan string)

	for {
		go ping(channel, rand.Intn(100)*10)
		print(<-channel, "\n")
		go pong(channel, rand.Intn(100)*10)
		print(<-channel, "\n")

	}
}
