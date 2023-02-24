package main

import (
	"fmt"
	"time"
)

func main() {
	
	ch := make(chan string)
	
	go func() {
		for {
			ch <- "ping"
			time.Sleep(1 * time.Second)
		}
	}()
	
	go func() {
		for {
			ch <- "pong"
			time.Sleep(1 * time.Second)
		}
	}()

	for data := range(ch) {
		fmt.Println(data)
	}

}