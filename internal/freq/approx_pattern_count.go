package freq

import "github.com/tyto-go/internal/distance"

// Insert your ApproximatePatternCount function here, along with any subroutines you need
func ApproximatePatternCount(Text string, Pattern string, d int) int {
	T, P := len(Text), len(Pattern)
	count := 0
	for i := 0; i <= T-P; i++ {
		if distance.HammingDistance(Text[i:i+P], Pattern) <= d {
			count++
		}
	}
	return count
}
