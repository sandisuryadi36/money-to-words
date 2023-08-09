package functions

import (
	"strings"
)

var maxIndex = len(idUnits) - 1

/*
NumbersToWord convert numbers string to words with ID or EN language.
Accept decimal number with dot "." separator
*/
func NumbersToWords(numbers string, lang string) string {
	res := ""
	switch strings.ToLower(lang) {
	case "id":
		integer, decimal := processDecimal(numbers)
		separator := ""
		if decimal != "" {
			separator = " " + idDecimalSeparatorWord + " "
		}
		res = numbersToWordsID(integer) + separator + decimalNumbersToWords(decimal, "id")
	case "en":
		integer, decimal := processDecimal(numbers)
		separator := ""
		if decimal != "" {
			separator = " " + enDecimalSeparatorWord + " "
		}
		res = numbersToWordsEN(integer) + separator + decimalNumbersToWords(decimal, "en")
	}
	return res
}

func processDecimal(numbers string) (integer, decimal string) {
	seperates := strings.Split(numbers, decimalSeparator)
	if len(seperates) > 1 {
		integer = seperates[0]
		decimal = seperates[1]
		return
	}
	integer = seperates[0]
	decimal = ""
	return
}

func digitToUnit(digit int, lang string) string {
	unitsArray := [22]string{}
	switch strings.ToLower(lang) {
	case "id":
		unitsArray = idUnits
	case "en":
		unitsArray = enUnits
	}

	curIndex := digit / 3
	if curIndex <= maxIndex {
		return unitsArray[curIndex]
	}
	return unitsArray[maxIndex]
}

func numberToText(index int, lang string) string {
	numbersArray := [10]string{}
	switch strings.ToLower(lang) {
	case "id":
		numbersArray = idNumbers
	case "en":
		numbersArray = enNumbers
	}

	if index >= 0 && index < len(idNumbers) {
		return numbersArray[index]
	}
	return ""
}

func decimalNumbersToWords(numbers string, lang string) string {
	words := []string{}
	zero := "nol"
	switch strings.ToLower(lang) {
	case "id":
		zero = idZero
	case "en":
		zero = enZero
	}

	for _, char := range numbers {
		digit := int(char - '0')
		if digit == 0 {
			words = append(words, zero)
		} else {
			words = append(words, numberToText(digit, lang))
		}
	}
	return strings.Join(words, " ")
}

func numbersToWordsID(numbers string) string {
	numbersLength := len(numbers)
	numbersMaxIndex := numbersLength - 1

	// Zero number
	if numbersMaxIndex == 0 && numbers[0] == '0' {
		return idZero
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
					result += space + numberToText(currentNumber, "id") + " " + digitToUnit(digitCount, "id")
				} else if modGroup == 2 {
					/* hundred number */
					if currentNumber == 1 {
						result += space + "seratus"
					} else {
						result += space + numberToText(currentNumber, "id") + " ratus"
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
								result += " " + digitToUnit(digitCount-1, "id")
							}
						} else {
							if nextNumber == 1 {
								result += space + "sebelas"
							} else {
								result += space + numberToText(nextNumber, "id") + " belas"
							}
							/* Proses Next current number */
							if digitCount != 1 {
								result += " " + digitToUnit(digitCount-1, "id")
							}
						}
					} else {
						/* tens number */
						result += space + numberToText(currentNumber, "id") + " puluh"
					}
				}
			} else {
				/* unit number not zero */
				if modGroup == 0 && (int(numbers[i-2]-'0') != 0 || int(numbers[i-1]-'0') != 0) && digitCount != 0 {
					result += " " + digitToUnit(digitCount, "id")
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

func numbersToWordsEN(numbers string) string {
	numbersLength := len(numbers)
	numbersMaxIndex := numbersLength - 1

	// Zero number
	if numbersMaxIndex == 0 && numbers[0] == '0' {
		return enZero
	}

	var space string
	var result string

	i := 0
	for i != numbersLength {

		digitCount := numbersMaxIndex - i
		modGroup := digitCount % 3 // [2,1,0]
		currentNumber := int(numbers[i] - '0')

		if currentNumber != 0 {
			if modGroup == 0 {
				/* unit number not zero */
				result += space + numberToText(currentNumber, "en") + " " + digitToUnit(digitCount, "en")
			} else if modGroup == 2 {
				/* hundred number */
				result += space + numberToText(currentNumber, "en") + " hundred"
			} else {
				/* 10 - 19 number */
				if currentNumber == 1 {
					i++ // Skip Next number
					nextNumber := int(numbers[i] - '0')
					if nextNumber == 0 {
						result += space + "ten"
						/* Proses Next current number */
						if digitCount != 1 && (int(numbers[i-1]-'0') != 0 || int(numbers[i]-'0') != 0) {
							result += " " + digitToUnit(digitCount-1, "en")
						}
					} else {
						result += space + enDozen[nextNumber]
						/* Proses Next current number */
						if digitCount != 1 {
							result += " " + digitToUnit(digitCount-1, "en")
						}
					}
				} else {
					/* tens number */
					result += space + enTens[currentNumber]
				}
			}
		} else {
			/* unit number not zero */
			if modGroup == 0 && (int(numbers[i-2]-'0') != 0 || int(numbers[i-1]-'0') != 0) && digitCount != 0 {
				result += " " + digitToUnit(digitCount, "en")
			}
		}

		if i <= 1 {
			space = " "
		}
		i++
	}

	return strings.TrimSpace(result)
}
