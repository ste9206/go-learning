package main

import "fmt"

func main() {
	var a any // interface{}

	a = 7
	fmt.Println("a:", a)

	a = "Hi"

	fmt.Println("a", a)

	/* Rule of thumb: don't use any :)

	 Exceptions:
	 - Serialization
	 - Printing
	*/

	s := a.(string) // type assertion

	fmt.Println("s:", s)

	i := a.(int) // will panic
	
	// AKA: "comma, ok", "ok" is a bool
	x, ok :=a.(int)

	if ok {
		fmt.Println("x:", x)
	} else {
		fmt.Println("not an int (%T)\n", x)
	}
}