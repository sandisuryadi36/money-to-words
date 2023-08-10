
# Numbers To Words
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)


Convert numbers to words in multiple languages.

## Overview

`NumberWordsConverter` is a lightweight Go package that provides a simple and efficient way to convert numeric strings into their word representations. This package supports both Indonesian (`id`) and English (`en`) languages, allowing you to easily transform decimal numbers into words.

## Installation

To use the `NumberWordsConverter` library in your Go project, follow these steps:

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
	"github.com/yourusername/numberwordsconverter/convert"
)

func main() {
	number := "12345.67"
	language := "en"

	words := convert.NumbersToWords(number, language)
	fmt.Println(words) // Output: "twelve thousand three hundred forty-five point sixty-seven"
}

```
## Features
- Convert decimal numbers to words in Indonesian (id) or English (en) languages.
- Handle various numeric formats, including thousands separators and decimal points.
- Customize decimal separator and unit names based on the chosen language.

## API Reference
`convert.NumbersToWords(numbers string, lang string) string` Converts a numeric string into its word representation in the specified language.

- numbers: The numeric string to convert.
- lang: The language for word conversion ("id" for Indonesian or "en" for English).
- Returns: The word representation of the numeric string.

This project is licensed under the MIT License - see the LICENSE file for details.

Enjoy converting numbers to words with ease using the NumberWordsConverter library! If you encounter any issues or have suggestions, please feel free to open an issue or submit a pull request.

## Feedback

If you have any feedback, please reach out to us at sandisuryadi.mail@gmail.com