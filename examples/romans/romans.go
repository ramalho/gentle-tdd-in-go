package romans

type NumeralValue struct {
	numeral string
	value   int
}

var numeralValues = []NumeralValue{
	{"I", 1},
	{"IV", 4},
}

func RomanValue(numeral string) int {
	value := 0
	for len(numeral) > 0 {
		for _, nv := range numeralValues {
			if nv.numeral == numeral {
				value += nv.value

			}
		}
	}
	return 0
}
