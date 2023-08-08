package constant

const (
	DecimalSeparator = "." 

	IDZero = "nol"
)

var (
	IDTerbilangSatuan = [...]string{
		"", "satu", "dua", "tiga", "empat", "lima", "enam", "tujuh", "delapan", "sembilan",
	}

	IDTerbilangAwalanSatu = [...]string{
		"", "sepuluh", "seratus", "seribu", "sepuluh ribu", "seratus ribu", "satu juta", "sepuluh juta", "seratus juta", "satu miliyar", "sepuluh miliyar", "seratus miliyar", "satu triliun", "sepuluh triliun", "seratus triliun",
	}

	IDTerbilangAkhiran = [...]string{
		"", "puluh", "ratus", "ribu", "puluh ribu", "ratus ribu", "juta", "puluh juta", "ratus juta", "miliyar", "ratus miliyar", "triliun",
	}
)