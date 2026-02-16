package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"sync"
	"time"
)

func main() {
	urls := []string{
		"https://go.dev",
		"https://ardanlabs.com",
		"https://ibm.com/no/such/page",
	}

	// non concurrent
	start := time.Now()

	for _, url := range urls {
		stat, err := urlCheck(url)

		fmt.Printf("%q: %d (%v)\n", url, stat, err)
	}

	duration := time.Since(start)
	fmt.Printf("%d urls in %v\n", len(urls), duration)

	// concurrent with channels
	start = time.Now()

	fanOutResult(urls)
	duration = time.Since(start)
	fmt.Printf("%d urls in %v\n", len(urls), duration)

	// concurrent with wg
	start = time.Now()

	fanOutWait(urls)
	duration = time.Since(start)
	fmt.Printf("%d urls in %v\n", len(urls), duration)


	// if you need result -> channel
	// if you need to read only -> wg
}

func fanOutResult(urls []string) {
	type result struct {
		url string
		status int
		err error
	}

	ch := make (chan result)

	for _, url := range urls {
		go func() {
			r := result {url: url}

			// that because if there's a panic
			// it prevents the deadlock
			defer func() {
				ch <- r
			}()

			r.status, r.err = urlCheck(url)
		}()
	}

	for range urls {
		r := <- ch
		fmt.Printf("%q: %d (%v)\n", r.url, r.status, r.err)
	}

}

func urlCheck(url string)(int, error) {
	resp, err := http.Get(url)

	if err != nil {
		return 0, err
	}

	return resp.StatusCode, nil
}

func fanOutWait(urls []string) {
	var wg sync.WaitGroup

	type result struct {
		url string
		status int
		err error
	}

	wg.Add(len(urls))
	for _, url := range urls {
		go func() {
			defer wg.Done()
			urlLog(url)
		}()
	}
	// wait for goroutines to finish
	// if you need errors, check out errgroup
	wg.Wait()
}

func fanOutPool(urls []string) {
	var wg sync.WaitGroup
	ch := make (chan string)
	const size = 2

	// producer
	go func() {
		for _, url := range urls {
			ch <- url
		}
		close(ch)
	}()

	wg.Add(size)
	
	for range size {
		// consumers
		go func() {
			defer wg.Done()
			for url := range ch {
				urlLog(url)
			}
		}()
	}
	wg.Wait()
}




func urlLog(url string) {
	resp, err := http.Get(url)

	if err != nil {
		slog.Error("urlLog", "url", url, "error", err)
		return
	}

	slog.Info("urllog", "url", url, "status", resp.StatusCode)
}