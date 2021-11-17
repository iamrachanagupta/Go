package main

import "fmt"

func main() {
	s := make([]string, 0)
	fmt.Println("length of s:", len(s)) //outputs: length of s: 0
	s = append(s, "hello")
	fmt.Println("length of s:", len(s))    //outputs: length of s: 1
	fmt.Println("contents of s[0]:", s[0]) //outputs: contents of s[0]: hello
	s[0] = "goodbye"
	fmt.Println("contents of s[0]:", s[0]) //outputs: contents of s[0]: goodbye

	s2 := make([]string, 2)                  //index: 0, 1, 2
	fmt.Println("contents of s2[0]:", s2[0]) //outputs: contents of s2[0]:
	s2 = append(s2, "hello")                 //appends at the end of the slice
	fmt.Println("After appending the slice with a string ...")
	fmt.Println("contents of s2[0]:", s2[0]) //outputs: contents of s2[0]:
	fmt.Println("contents of s2[1]:", s2[1]) //outputs: contents of s2[0]:
	fmt.Println("contents of s2[2]:", s2[2]) //outputs: contents of s2[2]: hello
	fmt.Println("length of s2:", len(s2))    //outputs:  length of s2: 3
}
