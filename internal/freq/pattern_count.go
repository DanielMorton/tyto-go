package freq

func PatternCount(Text string, Pattern string) int {
	T, P := len(Text), len(Pattern)
	count := 0
	for i := 0; i <= T-P; i++ {
		if Text[i:i+P] == Pattern {
			count++
		}
	}
	return count
}
