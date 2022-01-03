package main

import (
	"fmt"
)

//notes
type person struct {
	first string	
	last string
	age int
}

type secretAgent struct {
	person
	ltk bool
}

//exercise 1:
type catperson struct {
	first string
	last string
	favFlavors[]string
}

//exercise 3:
type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main(){

	//notes
	//aggregates together, values of a diff type
	//example
	//similar to object or class, but not the terminology to be used in go
	//does not recommend just putting "James" and "Bond" and 32
	//recommends always using the field name, for clarity
	p1 := person {
		first: "James",
		last: "Bond",
		age: 32,

	}

	p2 := person {
		first: "Miss",
		last: "Moneypenny",
		age: 31,
	}

	

	//can access with dot notation
	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println(p1.first, p1.last)
	fmt.Println(p2.first, p2.last)

	//embed structs
	sa1 := secretAgent{
		person: person{
			first: "James",
			last: "Bond",
			age: 27,
		},
		ltk: true,
	}

	fmt.Println(sa1.first, sa1.last, sa1.age, sa1.ltk)
	//don't need to show that this was part of .person

	//if only need one usage of a struct an anonymous struct will do
	pa := struct {
		first string
		last string
		age int 
	}{
		first: "James",
		last: "Bond",
		age: 32,
	}

	fmt.Println(pa)

	//exercises:
	//exercise 1 cont
	catperson1 := catperson {
		first: "Chloe",
		last: "Catdog",
		favFlavors: []string{
			"catnip flavor", 
			"chicken flavor",
		},
	}

	catperson2 := catperson {
		first: "April",
		last: "Yolanda",
		favFlavors: {}string{
			"wheatgrass flavor", 
			"mint flavor",
		},
	}

	fmt.Println(catperson1.first, catperson1.last)

	for i, v := range catperson1.favFlavors {
		fmt.Println(i, v)
	}

	fmt.Println(catperson2.first, catperson2.last)

	for i, v := range catperson2.favFlavors {
		fmt.Println(i, v)
	}

	//exercise 2
	//map what has been done so far, in ex 1
	m := map[string]catperson{
		catperson1.last:catperson1,
		catperson2.last:catperson2,
	}

	fmt.Println(m)

	for k, v := range m {
		fmt.Println(k)
		fmt.Println(v.first)
		fmt.Println(v.last)
		for i, val := range v.favFlavors {
			fmt.Println(i, val)
		}
	}


	//exercise 3 cont
	t := truck {
		vehicle: vehicle{
			doors: 2,
			color: white,
		},
		fourWheel: true,
	}

	s := sedan {
		vehicle: vehicle{
			doors: 4,
			color: "black",
		},
		luxury: false,

	},

	fmt.Println(t)
	fmt.Println(s)
	fmt.Println(t.doors)
	fmt.Println(s.doors)


	//exercise 4
	//goal create and use an anonymous struct
	//anonymous struct
	a := struct{
		first string
		friends map[string]int 
		favDrinks []string 
	}{ 
		first: "James",
		friends: map[string]int{
			"Moneypenny": 555,
			"Q": 777,
			"M": 888,
		},
		favDrinks: []string{
			"Martini",
			"Water",
		},
	}
	fmt.Println(a.first)
	fmt.Println(a.friends)
	fmt.Println(a.favDrinks)
	
	for k,v : = range a.friends {
		fmt.Println(k, v)
	}

	for i, val := range a.favDrinks {
		fmt.Println(i, val)
	}










}