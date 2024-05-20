package distance

import "github.com/tyto-go/internal/dna"

func Neighbors(Pattern string, d int) []string {
	if d == 0 {
		return []string{Pattern}
	}
	if len(Pattern) == 1 {
		return dna.Dna
	}

	neighborSet := make(map[string]bool)
	suffixNeighbors := Neighbors(Pattern[1:], d)
	for _, sn := range suffixNeighbors {
		if HammingDistance(sn, Pattern[1:]) < d {
			for _, l := range dna.Dna {
				neighborSet[l+sn] = true
			}
		} else {
			neighborSet[Pattern[0:1]+sn] = true
		}
	}
	neighbors := make([]string, 0)
	for n := range neighborSet {
		neighbors = append(neighbors, n)
	}
	return neighbors
}
