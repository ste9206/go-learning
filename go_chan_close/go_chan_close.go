package main

import "fmt"

func main() {
	ch := make(chan int)

		go func() {
		for i := range 4 {
			ch <- i
		}
		close(ch)
	}()
	
	// iterating over channel, we don't know how many elements
	// are there -> close(ch) is the way to go
	for v := range ch {
		fmt.Println(">>", v)
	}

	v, ok := <- ch // ch is closed
	fmt.Println("closed:", v, "ok:", ok) // won't panic, and the 0 value of integers

	/*
		The "for range" above does

		for {
			v, ok := <- ch

			if !ok {
				break
			}

			fmt.Println(">>", v)
		}
	*/

	// var ch chan int // ch is nil
}