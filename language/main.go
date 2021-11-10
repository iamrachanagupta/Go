package main

import (
	"fmt"

	"github.com/imrachanagupta/GO/language/mapper"
)

func main() {
	fmt.Println(mapper.Greet("Howdy, what's new?"))
	fmt.Println(mapper.Greet("Comment allez vous?"))
	fmt.Println(mapper.Greet("Wie geht es Ihnen?"))
	fmt.Println(mapper.Greet("Guten morgen"))
}
