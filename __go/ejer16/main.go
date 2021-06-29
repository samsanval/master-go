package main

import (
	"fmt"
	"time"
)

func main() {
	channelOne := make(chan time.Duration)
	go loop(channelOne)
	fmt.Println("Here I come")

	msg := <-channelOne
	fmt.Println(msg)
}
func loop(channel chan time.Duration) {
	start := time.Now()
	for i := 0; i < 1000000000; i++ {

	}
	end := time.Now()
	channel <- end.Sub(start)
}
