package main

import (
	"fmt"
	"sort"
	"unsafe"
)

func main() {
	cart := []string{"apple", "orange", "banana"}

	fmt.Println("len", len(cart))
	fmt.Println("Cart1", cart[1])

	// index + value
	for i,c := range cart {
		fmt.Println(i, c)
	}

	// just value
	for _, ca := range cart {
		fmt.Println(ca)
	}

	cart = append(cart, "milk")
	fmt.Println(cart)

	// slicing operator, half-open

	fruit := cart[:3]

	fmt.Println("fruit:", fruit)

	fruit = append(fruit, "lemon")

	fmt.Println("fruit", fruit)
	// milk is replaced by lemon
	fmt.Println("cart", cart)


	// slice is a struct 

	type slice struct {
		array unsafe.Pointer
		len int
		cap int
	}

	// cap is the capacity -> if there is no more capacity, it should reallocate

	// empty slice with 
	s1 := make([]int, 10) // -> [0,0,0,0,0,0,0,0,0,0]
	s2 := s1[3:7]


	var s []int

	for i := range 10_000 {
		s = appendInt(s, i)
	}

	out := concat([]string { "A","B"}, []string{"C"})

	fmt.Println(out)


	values := [] float64{3,1,2}
	fmt.Println(median(values))

	type Player struct {
		Name string
		Score int
	}


	players := []Player {
		{"Rick", 10_000},
		{"Morty", 11},
	}

	// value semantic "for" loop - not changing the values
	for _, p := range players {
		p.Score += 100
	}

	fmt.Println(players)

	// "pointer" semantic
	for i := range players {
		players[i].Score += 100
	}
	fmt.Println(players)

}


func concat(s1, s2 []string) []string {
	s := make([]string, len(s1)+len(s2))
	copy (s, s1)

	//s[x:] -> da x in avanti
	// s[:x] -> prima di x
	copy(s[len(s1):], s2)

	return s
}

func median(values []float64) float64 {

	// va fatto perchè puntano alla stessa area di memoria
	sorted := make ([]float64, len(values))

	copy(sorted, values)
	sort.Float64s(sorted)

	i := len(sorted) /2

	if len(sorted) %2 == 1 {
		return sorted[i]
	}

	mid := (sorted[i-1] + sorted[i])/2

	return mid
}


func appendInt(s []int, v int) []int {
	i :=len(s)

	if len(s) == cap(s) {
		// no more space in underlying array
		// need to reallocate and copy

		size := 2 * (len(s)+1)
		fmt.Println(cap(s), "-->", size)
		ns := make([]int, size)

		copy(ns, s)

		s = ns[:len(s)]
	}

	s = s[:len(s) + 1]
	s[i] = v
	return s
}
