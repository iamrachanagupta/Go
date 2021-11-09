package main

import (
	"fmt"

	"github.com/imrachanagupta/GO/leftpad"
)

func main() {
	fmt.Println(leftpad.Format("Hello", 15))
	fmt.Println(leftpad.Format("Rachana", 15))
	fmt.Println(leftpad.Format("Internationalization", 15))
	fmt.Println(leftpad.FormatRune("Hello", 15, '😃'))
	fmt.Println(leftpad.FormatRune("Goodbye", 15, '😃'))
}
