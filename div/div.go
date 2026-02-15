package main

import "fmt"

// go rule: no panic!

func main() {
	fmt.Println(safeDiv(6, 3))
	fmt.Println(safeDiv(6, 0))
}

/* using named return values:
- defer/recover to change return error value
- documentation
*/

// in go i can name return variables
func safeDiv(a,b int) (q int, err error) {

	// a panic can be caught inside a defer
	// and inside a recover()
	defer func() {
		if e := recover(); e != nil {
			err := fmt.Errorf("%v", e)

		}
	}()



	return div(a,b), nil
}

func div(a,b int) int {
	// this is going to panic!
	return a/b
}