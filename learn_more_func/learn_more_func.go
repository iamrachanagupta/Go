package main

import "fmt"

func addone(a int) int {
	return a + 1
}

func addTwo(a int) int {
	return a + 2
}

//function passing as a parameter to another function
func printOperation(a int, f func(int) int) {
	fmt.Println(f(a))
}

//function that returns a function
func makeAdder(b int) func(int) int {
	return func(a int) int {
		return a + b
	}
}

//Putting all together, passing a function and returning a function
func makeDoubler(f func(int) int) func(int) int {
	return func(a int) int {
		b := f(a)
		return b * 2
	}
}

func main() {
	//Ex1: creating a reference to the function
	myAddone := addone
	fmt.Println(addone(10))   //outputs : 11
	fmt.Println(myAddone(10)) //outputs : 11

	//Ex2: defining function inside main function. This function is accessible only within main
	myfunc := func(a int) int {
		return a + 1
	}
	fmt.Println(myfunc(20)) //outputs: 21

	//Ex3: function passing to another function
	printOperation(1, addone) //outputs: 2
	printOperation(2, addTwo) //outputs: 4

	//Ex4: modifying local variable inside a function defined inside main
	b := 2
	myfunc2 := func(a int) {
		b = b + a
	}
	myfunc2(10)
	fmt.Println(b) //outputs: 12

	//Ex5: returning a function to a local variable
	myOne := makeAdder(1)
	myTwo := makeAdder(2)
	fmt.Println(myOne(15)) //outputs: 16
	fmt.Println(myTwo(25)) //outputs: 27

	//Ex6: Putting all together, passing a function and returning a function
	double_myOne := makeDoubler(myOne)
	fmt.Println(double_myOne(15)) //outputs: 32, 16*2 = 32

}
