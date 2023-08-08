package functions

import (
	"strconv"

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

func Terbilang(angka string) string {
	angkaLength := len(angka)
	angkaMaxIndex := angkaLength - 1

	// Angka Nol
	if angkaMaxIndex == 0 && angka[0] == '0' {
		return C.IDZero
	}

	var space string
	var result string

	i := 0
	for i != angkaLength {

		digitCount := angkaMaxIndex - i
		modGroup := digitCount % 3 // [2,1,0]
		curAngka := int(angka[i] - '0')

		if digitCount == 3 && curAngka == 1 && (i == 0 || (int(angka[i-2]-'0') == 0 && int(angka[i-1]-'0') == 0)) {
			/* Angka Seribu */
			result += space + "seribu"
		} else {
			if curAngka != 0 {
				if modGroup == 0 {
					/* Angka Satuan Bukan Nol */
					result += space + numberToText(curAngka) + digitToUnit(digitCount)
				} else if modGroup == 2 {
					/* Angka Ratusan */
					if curAngka == 1 {
						result += space + "seratus"
					} else {
						result += space + numberToText(curAngka) + " ratus"
					}
				} else {
					/* Angka Sepuluh dan Belasan */
					if curAngka == 1 {
						i++ // Skip Next Angka
						nextAngka := int(angka[i] - '0')
						if nextAngka == 0 {
							result += space + "sepuluh"
							/* Proses Next Angka Sekarang */
							if digitCount != 1 && (int(angka[i-2]-'0') != 0 || int(angka[i-1]-'0') != 0) {
								result += " " + digitToUnit(digitCount-1)
							}
						} else {
							if nextAngka == 1 {
								result += space + "sebelas"
							} else {
								result += space + numberToText(nextAngka) + " belas"
							}
							/* Proses Next Angka Sekarang */
							if digitCount != 1 {
								result += " " + digitToUnit(digitCount-1)
							}
						}
					} else {
						/* Angka Puluhan */
						result += space + numberToText(curAngka) + " puluh"
					}
				}
			} else {
				/* Angka Satuan Nol */
				if modGroup == 0 && (int(angka[i-2]-'0') != 0 || int(angka[i-1]-'0') != 0) && digitCount != 0 {
					result += " " + digitToUnit(digitCount)
				}
			}
		}

		if i <= 1 {
			space = " "
		}
		i++
	}

	return result
}

func sToi(val string) uint64 {
	res, _ := strconv.ParseUint(val, 10, 64)
	return res
}
