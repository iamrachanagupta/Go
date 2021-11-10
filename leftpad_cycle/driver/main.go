package main

import (
	"fmt"

	leftpad "github.com/imrachanagupta/GO/leftpad_cycle"
)

func main() {
	fmt.Println(leftpad.Format("Hello", 15))
	fmt.Println(leftpad.Format("Rachana", 15))
	fmt.Println(leftpad.Format("Internationalization", 15))
	fmt.Println(leftpad.FormatRune("Hello", 15, '😃'))
	fmt.Println(leftpad.FormatRune("Goodbye", 15, '😃'))
	fmt.Println(leftpad.FormatRune("RachanaGupta", 15, '😃'))
}
