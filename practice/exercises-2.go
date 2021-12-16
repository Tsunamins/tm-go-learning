package main

import (
	"fmt"
)

// exercise 3
// goal: create typed and untyped consts
const (
	a3 = 42 //untyped
	b3 int = 43 //typed
)

//exercise 6
//iota usage with years as example
const (
	a6 = 2017 + iota
	b6 = 2017 + iota
	c6 = 2017 + iota
	d6 = 2017 + iota
)



// update later to only use one main throughout exercises
func main() {

	//exercise 1
	//goal: print number in decimal, binary, hex
	a := 42
	fmt.Printf("%d\t%b\t%#x", a, a, a)
	// d = decimal (i.e. 42); b = binary; #x = hexidecimal

	//exercise 2
	//goal is to assign the value of a check such as == or != to a variable, then print out the t/f values
	a2 := (42 == 42)
	b2 := (42 <= 43)
	c2 := (42 >= 43)
	d2 := (42 != 43)
	e2 := (42 < 43)
	f2 := (42 > 43)
	fmt.Println(a2, b2, c2, d2, e2, f2)

	//exercise 3 cont
	fmt.Println(a3, b3)

	//exercise 4
	//goal: assign int to variable; print int in decimal, binary, hex; shift bits of int over 1 positio to left and assign that to variable;
	// then print last created to dec, bin, hex

	a4 := 42
	fmt.Printf("%d\t%b\t%#x", a4, a4, a4)
	b4 := a4 << 1
	fmt.Printf("%d\t%b\t%#x", b4, b4, b4)

	//exercise 5
	//goal just for a raw string literal example
	a5 := `here is a string
	as
	an
	example of a
	raw "string literal"`
	fmt.Println(a5)

	//exercise 6 cont
	//iota usage
	fmt.Println(a6)
	fmt.Println(b6)
	fmt.Println(c6)
	fmt.Println(d6)













}