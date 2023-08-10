
# Numbers To Words
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/sandisuryadi36/number-to-words/blob/master/LICENSE) [![Go Reference](https://pkg.go.dev/badge/github.com/golang/go.svg)](https://pkg.go.dev/github.com/sandisuryadi36/number-to-words)


Convert numbers to words in multiple languages.

## Overview

`NumberToWordsConverter` is a lightweight Go package that provides a simple and efficient way to convert decimal number into their word representations. This package supports both Indonesian (`id`) and English (`en`) languages, allowing you to easily transform decimal numbers into words.

## Installation

To use the `NumberToWordsConverter` library in your Go project, follow these steps:

1. Install the package using `go get`:
```bash
go get github.com/sandisuryadi36/number-to-words
```
2. Import the package in your Go code:
```go
import "github.com/sandisuryadi36/number-to-words/convert"
```

## Usage/Examples

```go
package main

import (
	"fmt"
	"github.com/sandisuryadi36/number-to-words/convert"
)

func main() {
	number := 12345.67
	language := "en"

	words := convert.NumberToWords(number, language)
	fmt.Println(words) // Output: "twelve thousand three hundred forty five point sixty-seven"
}

```
## Features
- Convert decimal numbers to words in Indonesian (id) or English (en) languages.
- Handle decimal points.
- Handle hight range number

## Function Reference
`convert.NumberToWords(numbers float64, lang string) string`

Converts a decimal number into its word representation in the specified language.

- numbers: The decimal number to convert.
- lang: The language for word conversion ("id" for Indonesian or "en" for English).
- Returns: The word representation of the decimal number.

This project is licensed under the MIT License - see the LICENSE file for details.

Enjoy converting numbers to words with ease using the NumberWordsConverter library! If you encounter any issues or have suggestions, please feel free to open an issue or submit a pull request.

## Feedback

If you have any feedback, please reach out to us at sandisuryadi.mail@gmail.com