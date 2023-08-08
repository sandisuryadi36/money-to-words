package main

import (
	"fmt"

	"github.com/sandisuryadi36/money-to-words/functions"
)

func main() {
	v := float64(10000)
	fmt.Println(functions.ConvertMoneyToWords(v))
}