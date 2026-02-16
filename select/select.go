package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Buffered channel to avoid goroutine leak
	// Uber go guidelines: channel size is always unbuffered or a buffer of 1
	ch1, ch2 := make(chan string, 1), make (chan string, 1)

	go func() {
		time.Sleep(100 * time.Millisecond)
		ch1 <- "one"
	}()
	
	go func() {
		time.Sleep(200 * time.Millisecond)
		ch2 <- "two"
	}()

	cxt, cancel := context.WithTimeout(context.Background(), 10 * time.Millisecond)
	defer cancel()
	
	/*
		go routine leak:
		- context timeout runs first
		- at 100 ms, go routine 1 want to send "one" and gets stuck
		- nobody is going to listen on it anymore
		- in a server, it can hold resources
		easy solution: buffered channel (add 1 in the top)
	*/

	select {
		case v := <- ch1:
			fmt.Println("ch1:", v)
		case v:= <- ch2:
			fmt.Println("ch2:", v)
		/* case <- time.After(10 * time.Millisecond):
			fmt.Println("timeout") */
		case <- cxt.Done():
			fmt.Println("context timeout")
	}		

}