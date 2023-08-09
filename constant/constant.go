package constant

const (
	DecimalSeparator = "." 
	
	IDDecimalSeparatorWord = "koma" 
	IDZero = "nol"

	ENDecimalSeparatorWord = "point" 
	ENZero = "zero"
)

var (
	IDNumbers = [...]string{
		"", "satu", "dua", "tiga", "empat", "lima", "enam", "tujuh", "delapan", "sembilan",
	}

	IDUnits = [...]string{
		"", "ribu", "juta", "milyar", "triliun", "quadriliun", "quintiliun", "sextiliun", "septiliun", "oktiliun", "noniliun", "desiliun", "undesiliun", "duodesiliun", "tredesiliun", "quattuordesiliun", "quindesiliun", "sexdesiliun", "septendesiliun", "oktodesiliun", "novemdesiliun", "vigintiliun",
	}
	
	ENNumbers = [...]string{
		"", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
	}

	ENDozen = [...]string{
		"", "eleven", "twelfe", "thirdteen", "fourteen", "fifteen", "sixteen", "seventeen", "eightteen", "nineteen",
	}

	ENTens = [...]string{
		"", "ten", "twenty", "thirdty", "fourty", "fifty", "sixty", "seventy", "eightty", "ninety",
	}

	ENUnits = [...]string{
		"", "thousand", "million", "billion", "trillion", "quadrillion", "quintillion", "sextillion", "septillion", "octillion", "nonillion", "decillion", "undecillion", "duodecillion", "tredecillion", "quattuordecillion", "quindecillion", "sexdecillion", "septendecillion", "octodecillion", "novemdecillion", "vigintillion",
	}
	
)