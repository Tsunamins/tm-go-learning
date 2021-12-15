package main

import (
	"fmt"
)

// by 'at package level', he means outside of main
//exercise 2
 var x2 int
 var y2 string
 var z2 bool

//exercise 3, said to assign at pkg level also
var x3 int = 42
var y3 string = "James Bond"
var z3 bool = true

//exercise 4 was demonstrated also at pkg level
type catdog int 
var x4 catdog

//exercise 5
var y5 int


func main(){
	// exercise 1
	x := 42
	y := "James Bond"
	z := true

	//single line, then multi-line
	fmt.Println(x, y, z)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)


	//exercise 2 continued
	//fmt.Println(x2)
	fmt.Println(y2)
	fmt.Println(z2)
	//the values currently assigned are called "the zero value" (0, "", and false)

	//exercise 3 continued
	//instructions were to assign via short decl, with Sprintf formatting (\t are tabs)
	s := fmt.Sprintf("%v\t%v\t%v", x3, y3, z3)
	fmt.Println(s)

	//exercise 4 continued

	fmt.Println(x4)
	fmt.Printf("%T", x4)
	x4 = 42
	fmt.Println(x4)

	//exercise 5 continued
	y5 = int(x4)
	fmt.Println(y5)
	fmt.Printf("%T", y5)














}