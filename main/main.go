package main

import (
	"flag"
	"fmt"

	"github.com/sandisuryadi36/number-to-words/functions"
)

func main() {
	num := flag.Float64("n", 0, "# of iterations")
	flag.Parse()

	v := float64(*num)
	fmt.Println(functions.ConvertMoneyToWords(v))
}
