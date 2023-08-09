package functions

import (
	"strings"

	C "github.com/sandisuryadi36/number-to-words/constant"
)

var maxIndex = len(C.IDUnits) - 1

func digitToUnit(digit int) string {
	curIndex := digit / 3
	if curIndex <= maxIndex {
		return C.IDUnits[curIndex]
	}
	return C.IDUnits[maxIndex]
}

func numberToText(index int) string {
	if index >= 0 && index < len(C.IDNumbers) {
		return C.IDNumbers[index]
	}
	return ""
}

func NumberToWords(numbers string) string {
	numbersLength := len(numbers)
	numbersMaxIndex := numbersLength - 1

	// Zero number
	if numbersMaxIndex == 0 && numbers[0] == '0' {
		return C.IDZero
	}

	var space string
	var result string

	i := 0
	for i != numbersLength {

		digitCount := numbersMaxIndex - i
		modGroup := digitCount % 3 // [2,1,0]
		currentNumber := int(numbers[i] - '0')

		if digitCount == 3 && currentNumber == 1 && (i == 0 || (int(numbers[i-2]-'0') == 0 && int(numbers[i-1]-'0') == 0)) {
			/* one hundred */
			result += space + "seribu"
		} else {
			if currentNumber != 0 {
				if modGroup == 0 {
					/* unit number not zero */
					result += space + numberToText(currentNumber) + " " + digitToUnit(digitCount)
				} else if modGroup == 2 {
					/* hundred number */
					if currentNumber == 1 {
						result += space + "seratus"
					} else {
						result += space + numberToText(currentNumber) + " ratus"
					}
				} else {
					/* 10 - 19 number */
					if currentNumber == 1 {
						i++ // Skip Next number
						nextNumber := int(numbers[i] - '0')
						if nextNumber == 0 {
							result += space + "sepuluh"
							/* Proses Next current number */
							if digitCount != 1 && (int(numbers[i-1]-'0') != 0 || int(numbers[i]-'0') != 0) {
								result += " " + digitToUnit(digitCount-1)
							}
						} else {
							if nextNumber == 1 {
								result += space + "sebelas"
							} else {
								result += space + numberToText(nextNumber) + " belas"
							}
							/* Proses Next current number */
							if digitCount != 1 {
								result += " " + digitToUnit(digitCount-1)
							}
						}
					} else {
						/* tens number */
						result += space + numberToText(currentNumber) + " puluh"
					}
				}
			} else {
				/* unit number not zero */
				if modGroup == 0 && (int(numbers[i-2]-'0') != 0 || int(numbers[i-1]-'0') != 0) && digitCount != 0 {
					result += " " + digitToUnit(digitCount)
				}
			}
		}

		if i <= 1 {
			space = " "
		}
		i++
	}

	return strings.TrimSpace(result)
}
