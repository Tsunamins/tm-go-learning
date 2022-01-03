package main

import (
	"fmt"
)


func main(){
	//basic example of returning a string
	s1 := foo3()
	fmt.Println(s1)

	x := bar4()
	fmt.Printf("%T\n", x) //type is what is returned

	//to run - must assign to a variable since it is returning something
	//i := x()
	//fmt.Println(i)
	//instead of this can also just do
	//fmt.Println(x())
	//or instead of all of above
	fmt.Println(bar4()())


	//callback - passing func as argument, then what is referred to as functional programming
	//usual example
	ii := []int{2,3,4,5,6,7}
	s := sum1(ii...)
	fmt.Println("all numbers", s)

	//callback example - only want even numbers added
	s2 := even(sum1, ii...)
	fmt.Println("even numbers", s2)

	//closure - enclose variable - limit it's scope to only that function
	//w/i - just a code block - usage can be limited to just that code blocke
	//y could not be printed outside this code blocke
	{
		y := 42
		fmt.Println(y)
	}

	//closure more spec example - each will be a diff x
	a := incrementor()
	b := incrementor()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())

	//recursion example - func calls itself - using traditional factorial ex
	n := factorial(4)
	fmt.Println(n)

	
}

//basic returning a string
func foo3() string {
	//s := "Hello world"
	//return s
	//or can do 
	return "Hello world"
}

//example returning a function - the func bar4 will return a function with an int as it's return "func() int"
func bar4() func() int {
	return func() int{
		return 451
	}
}

//callback - passing func as argument
//first working with what is known, usual example
func sum1(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}

	return total

}

//callback example then as only wanting even numbers
//the callback function is this part "(f func(xi ...int) int" - a function that will take a variadic parameter and return and int
//"vi ...int", is a second parameter - the incoming numbers
func even(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v % 2 == 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)

}

//closure example, more specific use case
func incrementor() func() int {
	var x int
	return func() int {
		x++
		return x
	}

}

	//recursion example - func calls itself - using traditional factorial ex
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}