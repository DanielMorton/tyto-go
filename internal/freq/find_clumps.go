package freq

func FindClumps(Genome string, k, L, t int) []string {
	G := len(Genome)
	wordFreq := make(map[string]int)
	for i := 0; i <= L-k; i++ {
		wordFreq[Genome[i:i+k]]++
	}

	patternSet := make(map[string]bool)
	for word, freq := range wordFreq {
		if freq >= t {
			patternSet[word] = true
		}
	}
	left, right := 0, L+1
	for right <= G {
		wordFreq[Genome[right-k:right]]++
		wordFreq[Genome[left:left+k]]--
		if wordFreq[Genome[right-k:right]] >= t {
			patternSet[Genome[right-k:right]] = true
		}
		left++
		right++
	}
	patterns := make([]string, 0)
	for p := range patternSet {
		patterns = append(patterns, p)
	}
	return patterns
}
