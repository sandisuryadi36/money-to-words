package functions

import (
	"fmt"
	"strconv"
	"strings"

	C "github.com/sandisuryadi36/number-to-words/constant"
)

// ConvertMoneyToWords : convert from float64 money value to words
func ConvertMoneyToWords(value float64) string {
	result := ""
	valueInString := fmt.Sprintf("%2.f", value)
	SplittedValue := strings.Split(valueInString, C.DecimalSeparator)
	IntegerValue := strings.TrimSpace(SplittedValue[0])
	// DecimalValue := SplittedValue[1]

	converted := 0
	for i := len(IntegerValue) - 1; i >= 0; i-- {
		if getChar(IntegerValue, i) == "0" {
			if i == 0 && len(IntegerValue) == 1 {
				return C.IDZero
			}
			converted++
			continue
		}

		if getChar(IntegerValue, i) == "1" {
			if converted > 0 {
				result = C.IDTerbilangAwalanSatu[converted] + " " + result
				continue
			}
			result = C.IDTerbilangSatuan[1] + " " + result
			continue
		}
	}

	return strings.TrimSpace(result)
}

// function to get characters from strings using index
func getChar(str string, index int) string {
	return string([]rune(str)[index])
}

func cToi(val string) int {
	res, _ := strconv.Atoi(val)
	return res
}
