package main

import (
	"fmt"
)

func main() {

	//notes:
	//arrays:
	//have to specify the size in advance
	var x [5]int
	fmt.Println(x) //[0 0 0 0 0]
	//the type of the array includes the length
	//so the length of 5 here is part of they type of array
	//var x [6]int is a different type

	//assign at a position:
	x[3] = 42
	fmt.Println(x) //[0 0 0 42 0]

	//go docs go on to recommend the use of slices instead of arrays

	//slices

	//composite data types - struct type

	// create a composite literal via a slice, example - overall for arrays, slices, structs and maps
	// y := type{values}
	//~ described as a slice of int and it's values
	// in effect: a Slice allows you to group together values of the same type
	//further, expression that creates a new instance each time it is evaluated
	y := []int{4, 5, 7, 8, 42}
	fmt.Println(y)
	fmt.Println(len(y)) //get length of slice of int:
	fmt.Println(cap(y)) //covered soon
	fmt.Println(y[0])   //get at certain position

	//loop over values in slice
	for i, v := range y {
		//print out for index then value
		fmt.Println(i, v)

	}

	//slice the slice, get portions of the slice
	fmt.Println(y[1:])  //from element 1 through end
	fmt.Println(y[1:3]) //from element 1 up to but not including position 3, i.e. 5, 7

	//loop through without range, can cond be len(y)?
	for i := 0; i <= 5; i++ {
		fmt.Println(i, y[i])
	}

	//appending, deleting from a slice
	y = append(y, 20, 50, 25)
	fmt.Println(y)

	// 'unfurl' the z below, take them all out, place them in the y
	z := []int{234, 456, 678, 987}
	y = append(y, z...)
	fmt.Println(y)

	//deleting is slice of slice, then append
	y = append(y[:2], y[4:]...) //removes from element 2 - 3

	//slice - make notes
	//good for if you know the underlying size needed
	//structure is type, length, capacity
	a := make([]int, 10, 12)
	fmt.Println(a)
	fmt.Println(len(a)) //length, slice length of 10
	fmt.Println(cap(a)) //capacity, capacity of underlying array 12
	a[0] = 42
	a[9] = 999
	a = append(a, 3423)

	fmt.Println(a)
	fmt.Println(len(a)) //length, slice length of 10
	fmt.Println(cap(a)) //capacity, capacity of underlying array 12

	a = append(a, 3423)
	a = append(a, 3423)

	fmt.Println(a)
	fmt.Println(len(a)) //length, slice length of 10
	fmt.Println(cap(a)) //note difference in capacity after appending past 12 capacity, doubles size of underlying array to house the growing number of appends

	//multi-dimentional slice
	jb := []string{"James", "Bond", "Chocolate", "Martini"}
	fmt.Println(jb)
	mp := []string{"Miss", "Moneypenny", "Strawberry", "Hazelnut"}
	fmt.Println(mp)

	//then to be multidimensional - slice of a slice of a string
	xp := [][]string{jb, mp}
	fmt.Println(xp) //[[contents of jb slice][contents of mp slice]]

	//maps, key/value store
	//composite literal structure:
	//map[string]int - this is the whole portion of the type
	m := map[string]int{
		"James":      32,
		"Moneypenny": 27, //always use a trailing comma

	}
	fmt.Println(m)
	//access via the key
	fmt.Println(m["James"])
	//noteworthy - happen to enter a key that is not present - will return the zero value
	//means to check -
	v, ok := (m["April"])
	fmt.Println(ok)
	fmt.Println(v)

	if v, ok := m["April"]; ok {
		fmt.Println("THIS IS THE IF PRINT", v)
	}

	//add to the map
	m["todd"] = 33

	for k, v := range m {
		fmt.Println(k, v) //range through to print all, key and value (k, v)
	}

	//range through slice of int generally, often denoted by xi or ii
	xi := []int{4, 5, 6, 7, 8}
	for i, v := range xi {
		fmt.Println(i, v)
	}

	//delete entry from map
	delete(m, "todd")
	fmt.Println(m)
	//can delete something that doesn't exist, i.e won't change anything, but also won't throw error
	//can use check as above to ensure the key is present first

	//exercises:
	//exercise 1
	//goal use composite literal - create array hold 5 values type int
	//assign values to e index p, range over them and print them
	// use format printing to print out type of array
	x1 := [5]int{5, 4, 3, 2, 1}

	for i, v := range x1 {
		fmt.Println(i, v)

	}
	fmt.Printf("%T", x1)

	//exercise 2
	// goal - use a slice instead - using comp literal
	//assign 10 values - range over to print each and then print type of slice
	//note difference, when using comp literal, is no value in bracekts
	x2 := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}

	for i, v := range x2 {
		fmt.Println(i, v)

	}
	fmt.Printf("%T", x2)

	//exercise 3
	//slice the slice to isolate parts of slice, print out different portions of slice
	x3 := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	fmt.Println(x3[0:5])
	fmt.Println(x3[5:]) //up to and including the end
	fmt.Println(x3[2:7])
	fmt.Println(x3[1:6])

	//exercise 4 - append
	x4 := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	//append once
	x4 = append(x4, 52)
	fmt.Println(x4)
	//append multiples
	x4 = append(x4, 53, 54, 55)
	fmt.Println(x4)
	//then append another slice to the existing slice
	y4 := []int{56, 57, 58, 59, 60}
	fmt.Println(y4)
	x4 = append(x4, y4...)
	fmt.Println(x4)

	//exercise 5
	//goal - delete via appending and slicing
	x5 := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	x5 = append(x[:3], x5...) //up to and not including p 3
	fmt.Println(x5)

	//exercise 6
	//goal - create slice to store names of all US states
	//- use make and append for the goal of:
	//- underlying array should not be created more than once
	//- print length of slice, capacity
	//print out all values and index position

	x6 := make([]string, 50, 50)
	fmt.Println(x6)
	fmt.Println(len(x6))
	fmt.Println(cap(x6))
	// do not do this followed by x6 = []string{"","",""} - causes wrong capacity
	//instead
	states := []string{` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`}

	for i, v := range states {
		x6[i] = v
	}

	for i := 0; i < len(x6); i++ {
		fmt.Println(x6[i])
	}

	fmt.Println(len(x6))
	fmt.Println(cap(x6))

	//exercise 7 - multi-dimensional, slice of slice
	x7 := []string{"James", "Bond", "Shaken not stirred"}
	y7 := []string{"Miss", "Moneypenny", "Helloooo James"}

	xxs := [][]string{x7, y7}
	fmt.Println(xxs)

	for i, xs := range xxs {
		fmt.Println("record: ", i)
		for j, val := range xs {
			fmt.Printf("\tindex position: %v \t value: \t %v \n", j, val)
		}
	}


	//exercise 8
	m8 := map[string][]string{
		`bond_james`: []string{`Shaken not stirred`, `martinis`, `women`},
		`moneypenny`: []string{`James Bond`, `Literature`, `Computer science`},
		`no_dr`: []string{`being evil`, `ice cream`, `sunsets`},
	}

	for k, v := range m8 {
		fmt.Println("This is the record for: ", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}
	

	//exercise 9
	//add to map from 8
	m8["fleming_ian"] = []string{`steaks`, `cigars`, `espionage`}

	for k, v := range m8 {
		fmt.Println("This is the record for: ", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}


	//exercise 10
	// remove from map
	delete(m8, `no_dr`)
	for k, v := range m8 {
		fmt.Println("This is the record for: ", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}



}
