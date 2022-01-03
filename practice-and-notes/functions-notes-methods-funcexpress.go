package main

import (
	"fmt"
)

type person1 struct {
	first string
	last string

}

type secretAgent1 struct {
	person1
	ltk bool
}

//to note once again the func structure is
//func (r receiver) identifier(parameters) (return(s)) { code } 

//methods
// so in next notes will create the function 'speak' --> func speak() - but will use an r receiver
//going to create function speak and attach it to anybody that is a type of secret agent (secretAgent1) --> thus a method
//this function speak() is now attached to this type - secretAgent1
//any value of that type/this type has access to this method
//i.e. below in main 'sa1' - has access to this speak() method
func (s secretAgent1) speak(){
	fmt.Println("I am", s.first, s.last, " - the secretAgent speak")

}

func (p person1) speak() {
	fmt.Println("I am", p.first, p.last, " - the person1 speak")
}

//interfaces
//allow us to define behavior
//in this example - any type that has the method speak(), is also of they type human - so sa1 is of type secretAgent and human
//idea of that is that a value can be of more than one type
type human interface{
	speak()
}

//so here in function bar1 - either a person or a secret agent can be passed into bar1 - as they both are a type of human per the human interface with the function speak
func bar1(h human){
	fmt.Println("I was passed into bar1(), I called human,", h)

}

func main(){

	sa1 := secretAgent1{
		person1: person1{
			"James",
			"Bond",
		},
		ltk: true,
	}
	fmt.Println(sa1)

	sa2 := secretAgent1{
		person1: person1{
			"Money",
			"Penny",
		},
		ltk: true,
	}

	p1 := person1{
		first: "Dr",
		last: "No",
	}
	fmt.Println(p1)

	// now the method speak will be called on sa1 - it has access to because they type secretAgent1 was attached to the method speak()
	sa1.speak()
	sa2.speak()

	// example to illustrate interface and polymorphism:
	bar1(sa1)
	bar1(sa2)
	bar1(p1)


	//anonymous notes:
	//anon ex - not anonymous - has a name
	foo()

	//this is anonymous - has to have parenthesis right after as well - note (){}()
	func(){
		fmt.Println("Anonymous func ran")

	}()

	//can also have params:
	func(x int){
		fmt.Println("The meaning of life", x)
	}(42)

	//function expression - will assign to a variable and then run below
	f := func(){
		fmt.Println("my first func expression")
	}
	f()

	//can do as well with params/args
	f2 := func(x int){
		fmt.Println("the year big brother started watching")
	}
	f2(1984)


}

	//anonymous function - notes - this is not anonymous - has an identifier foo1
	func foo1() {
		fmt.Println("foo ran")
	}