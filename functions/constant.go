package functions

const (
	decimalSeparator = "." 
	
	idDecimalSeparatorWord = "koma" 
	idZero = "nol"

	enDecimalSeparatorWord = "point" 
	enZero = "zero"
)

var (
	idNumbers = [...]string{
		"", "satu", "dua", "tiga", "empat", "lima", "enam", "tujuh", "delapan", "sembilan",
	}

	idUnits = [...]string{
		"", "ribu", "juta", "milyar", "triliun", "quadriliun", "quintiliun", "sextiliun", "septiliun", "oktiliun", "noniliun", "desiliun", "undesiliun", "duodesiliun", "tredesiliun", "quattuordesiliun", "quindesiliun", "sexdesiliun", "septendesiliun", "oktodesiliun", "novemdesiliun", "vigintiliun",
	}
	
	enNumbers = [...]string{
		"", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
	}

	enDozen = [...]string{
		"", "eleven", "twelfe", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eightteen", "nineteen",
	}

	enTens = [...]string{
		"", "ten", "twenty", "thirty", "fourty", "fifty", "sixty", "seventy", "eightty", "ninety",
	}

	enUnits = [...]string{
		"", "thousand", "million", "billion", "trillion", "quadrillion", "quintillion", "sextillion", "septillion", "octillion", "nonillion", "decillion", "undecillion", "duodecillion", "tredecillion", "quattuordecillion", "quindecillion", "sexdecillion", "septendecillion", "octodecillion", "novemdecillion", "vigintillion",
	}
	
)