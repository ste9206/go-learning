package main

import (
	"fmt"
	"time"
)

/* func main() {
	go fmt.Println("go routine")

	fmt.Println("main")

	for i := range 3 {
		// prior to Go 1.22 this was a bug
		go func() {
			fmt.Println("goroutine:", i)
		}()
	}

	time.Sleep(10 * time.Millisecond)
} */


func main() {

	ch := make(chan int)
	/* ch <- 7 // send
	v := <- ch // receive
	
	fmt.Println(v) // PANIC, because of deadlock
	*/

	go func() {
		ch <- 7
	}()

	v := <- ch
	fmt.Println(v)

	fmt.Println(sleepSort([]int{20, 30, 10})) // [10 20 30]

}

func sleepSort(values []int) {
	ch := make(chan int)

	for _, n := values {
		go func() {
			time.Sleep(time.Duration(n) * time.Microsecond)
			ch <- n
		}()
	}

	var out []int
	
	for range values {
		n := <- ch
		out = append(out, n)
	}
	return out
}

/* channel semantics
- send/receive to/from a channel will block until opposite operation(*)
- guarantee of delivery (receiving happens before sending): once you're done with your send
  you know that on the other side someone got a message
  - buffered channel has "n" non blocking sends (lose guarantee)
  - receive from a closed channel will return zero value without blocking
- use "comma ok" to check if channel was closed
- send to a close channel will panic
- channel "ownership": the owner of the channel (who sends) is the one closing it
- closing a closed or nil channel will panic
- send/receive to a nil channel will block forever
*/