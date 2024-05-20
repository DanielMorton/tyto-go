package freq

import "github.com/tyto-go/internal/distance"

func ApproximatePatternMatching(Pattern string, Text string, d int) []int {
	P := len(Pattern)
	T := len(Text)
	match := make([]int, 0)
	for i := 0; i <= T-P; i++ {
		if distance.HammingDistance(Text[i:i+P], Pattern) <= d {
			match = append(match, i)
		}
	}
	return match
}
