package main

import (
	"flag"
	"fmt"

	"github.com/sandisuryadi36/number-to-words/functions"
)

func main() {
	v := flag.String("n", "0", "# of iterations")
	lang := flag.String("l", "id", "# of iterations")
	flag.Parse()

	fmt.Println(functions.NumbersToWords(*v, *lang))
}
