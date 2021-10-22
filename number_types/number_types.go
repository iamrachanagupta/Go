package main

import "fmt"

func main() {
	/*
		aliases for the types:
		"byte" same as uint8
		"int" same as int32 or int64 based on OS
		"uint" same as uint32 or uint64
	*/
	var i int8 = 20
	var f float32 = 5.6
	fmt.Println(float32(i) + f)  //outputs 25.6
	fmt.Println(i + int8(f))     //outputs 25
	fmt.Println(i + int8(f+1.9)) // outputs 27
	i = 127                      //cannot assign > 127 otherwise throws error: overflows int8
	fmt.Println(i)               //outputs 127
	var j int32 = 40
	fmt.Println(int32(i) + j) //type casting is necessary. Outputs 167
}
