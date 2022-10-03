package iteration

// Repeat returns character repeated times.
func Repeat(character string, times int) string {
	var repeated string

	for i := 0; i < times; i++ {
		repeated += character
	}

	return repeated
}
