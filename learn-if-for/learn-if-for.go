package main

import "fmt"

func main() {

	a := 10
	/* Invalid if conditions-- throws compilation error
	if 0 {}
	if a = 5 {}   ---> assignment instead of comparision.
	*/

	if c := a / 2; a > 5 {
		b := a / 2
		fmt.Println("a is bigger than 5")
		fmt.Println(a, b) // outputs : 10 5
		fmt.Println(c)    //outputs 5
	} else {
		fmt.Println("a is smaller than 5")
		//fmt.Println(b) --> outside the scope of variable b, throws undefined b
		fmt.Println(c) //outputs 5
	}
	//fmt.Println(b) --> outside the scope of variable b, throws undefined b
	//fmt.Println(b) --> outside the scope of variable c, throws undefined c

	a1 := 3
	for i := 0; i < 10; i++ {
		if i == a1 {
			continue
		}
		if i > 7 {
			break
		}
		fmt.Print(i, " ") //outputs: 0 1 2 4 5 6 7
		fmt.Print("\n")
	}

	//for statement like while statement
	i := 0
	for i < 5 {
		fmt.Print(i, " ") //outputs: 0 1 2 3 4
		i = i + 1
	}
	fmt.Print("\n")
	//another way to use for loop
	i = 0
	for {
		fmt.Print(i, " ") //outputs 0 1 2 3 4 0 1 2 3 4 5 6 7 8 9 10
		i++
		if i > 10 {
			break
		}
	}
	fmt.Print("\n")

	//to iterate over every char in a string
	s := "Hello"
	for k, v := range s {
		fmt.Println(k, v, string(v))
	}
	/* outputs:
	0 72 H
	1 101 e
	2 108 l
	3 108 l
	4 111 o
	*/

	//to iterate over every rune in a string
	s = "👋 🌍"
	for k, v := range s {
		fmt.Println(k, v, string(v))
	}
	/* outputs:
	0 128075 👋
	4 32
	5 127757 🌍

	Note that the offset is measured in bytes
	*/

}
