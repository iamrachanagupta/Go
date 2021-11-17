package main

import "fmt"

func main() {
	m2 := map[string]int{
		"a": 1,
		"b": 2,
		"c": 50, //every line must end with a comma even the last line
	}

	for k, v := range m2 { //k- key, v- value
		fmt.Println(k, v) /* values of map can be printed in any order, but in this run it printed in proper order but thats not necessary for every run
		outputs:
		a 1
		b 2
		c 50
		*/
	}

	fmt.Println("b in m2:", m2["b"]) //outputs: b in m2: 2
	delete(m2, "b")                  // this deletes the key-value pair
	fmt.Println("b in m2:", m2["b"]) //outputs: b in m2: 0
}
