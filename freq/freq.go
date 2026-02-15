package main

import (
	"bufio"
	"fmt"
	"maps"
	"os"
	"regexp"
	"slices"
	"sort"
	"strings"
)

// What are the N most common words in sherlock.txt

var wordRe = regexp.MustCompile(`[a-zA-Z]+`)

// Code that runs before main
// - var expressions
// - init function


func main() {
	// mapDemo()
	file, err := os.Open("sherlock.txt")

	if err != nil {
		fmt.Println("Error", err)
		return
	}

	defer file.Close()

	s := bufio.NewScanner(file)
	nLines := 0


	freq := make(map[string]int)

	for s.Scan() {
		nLines++
		words := wordRe.FindAllString(s.Text(), -1)
		for _, word := range words {
			freq[strings.ToLower(word)]++
		}
	}

	if err := s.Err(); err != nil {
		fmt.Println("Error", err)
		return	
	}

	top := topN(freq, 10)
	fmt.Println(top)
}

// topN returns the "n" most common words in freq.
func topN(freq map[string]int, n int) []string {
	words := slices.Collect(maps.Keys(freq))
	sort.Slice(words, func(i, j int)bool {
		wi, wj := words[i], words[j]
		// sort in reverse order
		return freq[wi] > freq[wj]
	})

	n = min(n, len(words))
	return words[:n]
}


func mapDemo() {
	heros := map[string]string {
		"Superman": "Clark",
		"Wonder Woman": "Diana",
		"Batman": "Bruce",
	}

	// keys
	for k := range heros {
		fmt.Println(k)
	}

	// keys, values
	for k, v := range heros {
		fmt.Println(v, "is", k)
	}

	// only values
	for _, v := range heros {
		fmt.Println(v)
	}

	n := heros["Batman"]
	fmt.Println(n)

	n = heros["Aquaman"]
	fmt.Println(n)

	// this to look if key exists
	n, ok := heros["Aquaman"]

	if ok {
		fmt.Printf("%q\n", n)
	} else {
		fmt.Println("Aquaman not found")
	}

	delete(heros, "Batman")
	fmt.Println(heros)

}