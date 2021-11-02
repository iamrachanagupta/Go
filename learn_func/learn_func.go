package main

import "fmt"

//function definition can be put anywhere even after main.
func addNumbers() {
	fmt.Println(2 + 3)
}

func main() {
	addNumbers()
	add(10, 20)
	add(-100, 30)
	a := subtract(100, 40)
	fmt.Println(a)

	add, sub, div := addSubDiv(100, 50)
	fmt.Println(add, sub, div)
	//Ignoring a return value: just use "_"
	ad, _, di := addSubDiv(100, 20)
	fmt.Println(ad, di)
}

func add(a int, b int) {
	fmt.Println(a + b)
}

func subtract(a int, b int) int {
	return a - b
}

//returning multiple parameters
func addSubDiv(a int, b int) (int, int, int) {
	return a + b, a - b, a / b
}

/* Note: func overloading, optional parameters, named parameters are not available in Go lang
All function calls are called by value in Go lang */
