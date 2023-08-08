package functions

import (
	"fmt"
	"strings"

	"github.com/sandisuryadi36/money-to-words/constant"
)

// ConvertMoneyToWords : convert from float64 money value to words
func ConvertMoneyToWords(value float64) string {
	result := ""
	valueInString := fmt.Sprintf("%2.f", value)
	SplittedValue := strings.Split(valueInString, constant.DecimalSeparator)
	IntegerValue := SplittedValue[0]
	// DecimalValue := SplittedValue[1]

	converted := 0
	for i:= len(IntegerValue)-1; i >= 0; i-- {
		if getChar(IntegerValue, i) == "0" {
			if i == 0 && len(IntegerValue) == 1 {
				return constant.IDZero
			}
			converted++
			continue
		}

		if getChar(IntegerValue, i) == "1" {
			if converted > 0 && converted <= 5 {
				result = constant.IDTerbilangAwalanSatu[converted] + " " + result
				continue
			}
			result = constant.IDTerbilangSatuan[i] + " " + result
			continue
		}
	}

	return strings.TrimSpace(result)
}

// function to get characters from strings using index
func getChar(str string, index int) string {
	return string([]rune(str)[index])
}