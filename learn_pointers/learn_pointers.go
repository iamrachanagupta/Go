package main

import "fmt"

func main() {
	//Ex1: Understanding pointers
	a := 10
	b := &a
	c := a
	fmt.Println(a, *b, c, &a, b, &c) //outputs: 10 10 10 0xc0000aa058 0xc0000aa058 0xc0000aa070
	a = 20
	c = 40
	fmt.Println(a, *b, c, &a, b, &c) //outputs: 20 20 40 0xc0000aa058 0xc0000aa058 0xc0000aa070
	*b = 30
	fmt.Println(a, *b, c, &a, b, &c) //outputs: 30 30 40 0xc0000aa058 0xc0000aa058 0xc0000aa070

	//Ex2: assigned as nil pointer by default
	var a_ptr *int
	fmt.Println(a_ptr) //outputs: <nil>
	//fmt.Println(*a_ptr) //outputs: panic: runtime error: invalid memory address or nil pointer dereference
	//Cannot read or write a nil pointer, otherwise it will throw a panic

	//Ex3: allocating memory of type int to the a_ptr
	a_ptr = new(int)
	fmt.Println(a_ptr)  //outputs: 0xc0000120e0
	fmt.Println(*a_ptr) //outputs: 0

	//Ex4: passing pointers to a function
	aa := 10
	fmt.Println(aa)
	setTo20(&aa)
	fmt.Println(aa)
}
func setTo20(bb *int) {
	*bb = 20
}
