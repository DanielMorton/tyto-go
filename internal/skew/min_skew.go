package skew

func MinimumSkew(Genome string) []int {
	skew := make([]int, 1)
	currSkew := 0
	minSkew := currSkew
	skew[0] = currSkew
	for _, c := range Genome {
		if c == 'C' {
			currSkew--
			if currSkew < minSkew {
				minSkew = currSkew
			}
		} else if c == 'G' {
			currSkew++
		}
		skew = append(skew, currSkew)
	}
	minSkewPos := make([]int, 0)
	for i, s := range skew {
		if s == minSkew {
			minSkewPos = append(minSkewPos, i)
		}
	}
	return minSkewPos
}
