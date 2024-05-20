package freq

func PatternMatching(Pattern string, Genome string) []int {
	matches := make([]int, 0)
	G := len(Genome)
	P := len(Pattern)
	for i := 0; i <= G-P; i++ {
		if Pattern == Genome[i:i+P] {
			matches = append(matches, i)
		}
	}
	return matches
}
