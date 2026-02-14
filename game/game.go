package main

import (
	"fmt"
	"slices"
)

func main() {
	var i Item
	// %#v prints the type
	// so good for debugging/logging
	fmt.Printf("i: %#v\n", i)

	a, b := 1, "1"

	fmt.Printf("a=%#v, b=%#v", a,b)

	i = Item{10,20} //must specify all the fields

	i = Item{X: 20, Y: 21} // i can even omit with that way

	fmt.Println((NewItem(10,20)))

	i.Move(10, 20)
	fmt.Printf("i move:%#v\n", i)

	p1 := Player{
		Name: "Paul",
	}

	// item is embedded
	fmt.Printf("p1:%#v\n", p1)
	fmt.Printf("p1.X:%#v\n", p1.X) // X is part of item, which is embedded
	
	p1.Move(100, 200)


	fmt.Println(p1.Found(Copper))
	fmt.Println(p1.Found(Key(7)))
	fmt.Println(p1.Keys)

	ms := [] Mover{
		&i, &p1,
	}

	moveAll(ms, 50, 70)

	for _, m := range ms {
		fmt.Println(m)
	}
}

// check go proverbs

type Item struct {
	X int
	Y int
}


/*func NewItem(x,y int) Item {

}

func NewItem(x,y int) *Item {

}
*/

// Value semantics: everyone has their own copy
// Pointer semantics: everyone share the same copy (heap, lock)
// pointer can be more expensive since it allocates it on the heap

func NewItem(x,y int) (*Item, error) {
	if x < 0 || x > maxX || y < 0 || y > maxY {
		return nil, fmt.Errorf("%d/%d of of bonds %d/%d", x, y, maxX, maxY)
	}

	i := Item{
		X:x,
		Y:y,
	}

	// go compiler does escape analysis and will allocate i on the heap
	return &i, nil
}

// value semantic
/*func NewItem(x,y int) (Item, error) {
	if x < 0 || x > maxX || y < 0 || y > maxY {
		return Item{}, fmt.Errorf("%d/%d of of bonds %d/%d", x, y, maxX, maxY)
	}
	
	i := Item{
		X:x,
		Y:y,
	}

	return i, nil
}*/

const (
	maxX = 600
	maxY = 400
)


/* "i" is called "the receiver"
 i is a pointer receiver

VALUE VS POINTER RECEIVER
 - in general use value semantics
 - try to keep same semantics on all methods

- when you must use pointer receiver:
1. you have a lock field (mutex, etc)
2. if you need to mutate the struct
3. decoding/unmarshalling

*/

func(i *Item) Move(dx, dy int) {
	i.X += dx
	i.Y += dy
}


type Player struct {
	Name string
	Item // Player embeds item
	Keys [] Key
}

type Key byte

const (
	Copper Key = iota + 1
	Jade
	Crystal
)



func (p *Player) Found(key Key) error {
	switch key {
	case Copper, Jade, Crystal:
			// ok
		default:
			return fmt.Errorf("unknown key: %q", key)
	
	}

	if !slices.Contains(p.Keys, key) {
		p.Keys = append(p.Keys, key)
	}
	
	return nil
}
/*

Interfaces
- set of methods (and types)
- we define interfaces as "what you need", not "what you provide"
	- Interfaces are small (stdlib average ~2 methods per interface)
	- If you have an interface with more than 4 methods, think again

	- Start with concrete types, discover interface
*/


func Sort(s Sortable) {

}

type Sortable interface {
	Less(i,j int) bool
	Swap(i,j int)
	Len() int
}

// Rule of thumb: Accept interface, return types

// go install stringer interface
// in ~/.zshrc
// export PATH="$(go env GOPATH)/bin:${PATH}"

func (k Key) String() string {
	switch k {
		case Copper:
			return "copper"
		
		case Jade:
			return "jade"
		
		case Crystal:
			return "crystal"

		default:
			return fmt.Sprintf("<Key %d>", k)
	}
}


// set of types and methods
type Mover interface {
	Move(int, int)
}

func moveAll(ms []Mover, dx, dy int) {
	for _,m := range ms {
		m.Move(dx, dy)
	}
}