package main

import "fmt"

func main() {
	var i int // declaration , automatically initialised to 0 by compiler
	j := 20   /* declaration with initialization without specifying type
	But it should be followed by its usuage like a print statement otherwise will throw error*/
	fmt.Println(i) //outputs 0
	fmt.Println(j) //outputs 20
	i = 15         //assignment
	fmt.Println(i) //outputs 15

}
