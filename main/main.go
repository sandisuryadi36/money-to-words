package main

import (
	"flag"
	"fmt"

	"github.com/sandisuryadi36/number-to-words/functions"
)

func main() {
	v := flag.String("n", "0", "# of iterations")
	flag.Parse()

	fmt.Println(functions.NumberToWords(*v))
}
