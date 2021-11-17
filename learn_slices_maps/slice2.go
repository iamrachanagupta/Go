package main

import "fmt"

func main() {
	s3 := []string{"a", "b", "c"}

	for k, v := range s3 {
		fmt.Println(k, v)
		/* outputs:
		0 a
		1 b
		2
		*/
	}

	s4 := s3[0:2]          //creating a slice, index 2 exclusive
	fmt.Println("s4:", s4) //outputs: s4: [a b]
	s3[0] = "d"
	fmt.Println("s4:", s4) //outputs: s4: [d b]

	var s5 []string
	s5 = s3
	s5[1] = "camel"
	fmt.Println("s3:", s3) //outputs: s3: [d camel c]

	modSlice(s3)
	fmt.Println("s3[0]:", s3[0]) //outputs: s3[0]: hello
}

func modSlice(s []string) {
	s[0] = "hello"
}
