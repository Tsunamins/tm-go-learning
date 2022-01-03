package main

import (
	"fmt"
)

func main(){
	foo()
	bar("James")
	s1 := woo("Moneypenny")
	fmt.Println(s1)
	//returns 2 vars so can assign accordingly
	x, y := mouse("Ian", "Fleming")
	//x and y are the returns of this function, not the arguments
	fmt.Println(x)
	fmt.Println(y)

	//variadic parameter example cont
	//can also pass in nothing i.e. foo2() or sum(), bc variadic is 0 to infinite amount
	//would need to be the final parameter/argument as it could be no values if there were multiple params/args 
	x := foo2(2,3,4,5,6,7)
	fmt.Println("The total is", x)

	// 'unfurl'  slice example
	xi := []int{2,3,4,5,6,7}
	//adds the ints above, xi..., he calls it 'unfurling' as in take all these and dump them into here
	x := sum(xi...)

	//defer example
	defer foo3()
	bar2()

}

//basic:
func foo(){
	fmt.Println("Hello from foo function")
}

//basic with 1 param
//note: everything in go is passed by value - what you see is what you get
func bar(s string){
	fmt.Println("Hello,", s)
}

//func assigned to variable, and will return a string:
// the func woo(with parameter) return-type-string
func woo(st string) string{
	return fmt.Sprint("Hello from woo, ", st)
}

//specify multiple params
// can do as follows, or to be more clear, fn string, ln string
// the func mouse(with two params of string) (will return a string and bool)
func mouse(fn, ln string) (string, bool) {
	a := fmt.Sprint(fn, ln, `, says "Hello`)
	b := true
	return a, b

}

//variadic parameter example:
// unlimited amt of params, ie ints in this case
//this becomes a slice of int, noted when printing out the type
//could pass in nothing all together so is 0 to infinite amount
func foo2(x ...int){
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sumoffoo2 := 0 //preferred, more readable, alternate is var sum int
	for i, v := range x {
		sumoffoo2 += v
		fmt.Println("for item in index position,", i, "we are now adding,", v, "to the total which is now", sumoffoo2) 

	}
	fmt.Println("The total is,", sumoffoo2)
	return sumoffoo2
}

//left off on unfurling a slice section
//using similar example as foo2
func sum(x...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("for item in index position", i, "we are now adding," v, "to the total which is now", sum)
	}
}

//deferring
func foo3(){
	fmt.Println("foo")


func bar2(){
	fmt.Println("bar")
}

