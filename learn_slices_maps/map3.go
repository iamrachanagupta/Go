package main

import "fmt"

func main() {
	m := map[string]int{
		"a": 1,
		"b": 2,
	}
	var m3 map[string]int

	fmt.Println("goodbye in m:", m["goodbye"]) //outputs: goodbye in m: 0
	//maps are reference types
	m3 = m
	m3["goodbye"] = 400
	fmt.Println("goodbye in m3:", m3["goodbye"]) //outputs: goodbye in m3: 400
	fmt.Println("goodbye in m:", m["goodbye"])   //outputs: goodbye in m: 400

	modMap(m)                                //passed as reference
	fmt.Println("cheese in m:", m["cheese"]) //outputs: cheese in m: 20
}

func modMap(m map[string]int) {
	m["cheese"] = 20
}
