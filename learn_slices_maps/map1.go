package main

import "fmt"

func main() {
	/* A map is just a data type that I sort of see it's a value of one data type called the key to a value of another data type.*/
	m := make(map[string]int) //key is of type string and value is of type int
	m["hello"] = 300
	h := m["hello"]
	fmt.Println("hello in m:", h)
	fmt.Println("a in m:", m["a"]) /*outputs: a in m: 0 (key "a" is non existent but still value is printed as 0 which is still a valid value for a int, hence this is not the way to know if a value exist corresponding to a key) */

	if v, ok := m["hello"]; ok { // ok is a boolean variable which is set to true only if the key-value pair exists
		fmt.Println("hello in m:", v) //outputs: hello in m: 300
	}
	if v, ok := m["a"]; ok {
		fmt.Println("a in m:", v) //This doesn't get printed as a is not a valid key
	}

	m["hello"] = 20
	fmt.Println("hello in m:", m["hello"]) //outputs: hello in m: 20
}
