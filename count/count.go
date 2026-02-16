package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

/*
	sync.atomic is lower level than mutex
*/

func mutexSolution() {
	// mutex should be on top of variable that is "guarding"
	var mu sync.Mutex
	count := 0

	nGR, nIter := 10, 1_000

	var wg sync.WaitGroup


	wg.Add(nGR)

	for range nGR {
		go func() {
			defer wg.Done()
			for range nIter {
				mu.Lock()
				count++
				mu.Unlock()
				/*
					fetch count
					increment count
					store count
				*/
				time.Sleep(time.Microsecond)
			}
		}()
	}
	wg.Wait()
	fmt.Println("count:", count)
}

/*
	use it only if you have a serious performance requirement
*/
func atomicSolution() {
	count := int64(0)

	nGR, nIter := 10, 1_000

	var wg sync.WaitGroup


	wg.Add(nGR)

	for range nGR {
		go func() {
			defer wg.Done()
			for range nIter {
				atomic.AddInt64(&count, 1)
				time.Sleep(time.Microsecond)
			}
		}()
	}
	wg.Wait()
	fmt.Println("count:", count)
}


func main() {
	mutexSolution()
	atomicSolution()
}

/*

go run -race ./count

"-race" is supported by:
-run
-build
-test

why not used everytime? runtime overhead

Rule of thumb: use "go test -race"

*/