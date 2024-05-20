package dna

func ReverseComplement(Pattern string) string {
	reverse := ""
	for i := len(Pattern) - 1; i >= 0; i-- {
		reverse += string(DnaComp[Pattern[i]])
	}
	return reverse
}
