package main

import (
	"fmt"
)

func main(){

	//examples nested for loop, for in general
	for i := 0; i <= 10; i++ {
		fmt.Printf("The outer loop: %d\n", i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t\t The inner loop: %d\n", j)
		}
	}

	//can be most basic as
	// for {}
	//like a while instead
	x := 1
	for x < 10 {
		fmt.Println(x)
		x++
	}
	fmt.Println("done.")

	//break example - with if
	y := 1
	for {
			if y > 9 {
				break
			}
			fmt.Println(y)
			y++
		
	}
	fmt.Println("done.")

	//continue example
	z := 0
	for {
		z++
		if z > 100 {
			break
		}

		if z%2 != 0 {
			continue
		}

	}
	fmt.Println(z)
	
	//else if
	a := 41
	if a == 40{
		fmt.Println(a)
	} else if a == 41 {
		fmt.Println(a)
	} else {
		fmt.Println("something else")
	}

	//switch
	//will go to first match, unless 'fallthrough', note that fallthrough will continue on to even false evaluations so 
	//best practice don't use fallthrough
	//example is when switch is not specified
	switch {
	case false:
		fmt.Println("is false")

	case (3 == 3):
		fmt.Println("is true")
		fallthrough

	case (5 == 5):
		fmt.Println("is a42lso true if fallthrough")
	
	default:
		fmt.Println("default available otherwise")
	}

	//exercise 1
	for i1 := 0; i1 <= 10000; i1++ {
		fmt.Println(i1)
	}

	//exercise2
	//goal print every rune code point of uppercase alphabet 3 times
	//explanation capital ascii's go from 65 through 90
	for i2 := 65; i2 <= 60; i2++ {
		fmt.Println(i2)

		for j2 := 0; j2 < 3; j2++{
		fmt.Printf("\t%#U\n", i2)
		}
	}

	//exercise3
	bd := 1985
	for bd < 2021 {
		fmt.Println(bd)
		bd++

	}

	//exercise 4
	bd4 := 1985
	for {
		if bd4 > 2021 {
			break
		}
		fmt.Println(bd4)
		bd++

	}

	//exercise 5
	for i5 := 10; i5 <= 100; i5++ {
		fmt.Printf("When %v is divided by 4 the remainder is %v\n", i5, i5%4)
	}



	//exercise 9
	//switch specified
	favSport := "surfing"
	switch favSport {
	case "skiing":
		fmt.Println("mountains")
	case "swimming":
		fmt.Println("calm ocean")

	case "surfing":
		fmt.Println("ocean with surf")

	case "wingsuit flying":
		fmt.Println("air")


	}







}