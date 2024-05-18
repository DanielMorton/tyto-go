package alignment

import "github.com/tyto-go/internal"

func EditDistance(s string, t string) int {
	scores := internal.MakeGrid[int](len(s)+1, len(t)+1)
	scores[0][0] = 0
	for i := 1; i <= len(s); i++ {
		scores[i][0] = scores[i-1][0] + 1
	}
	for j := 1; j <= len(t); j++ {
		scores[0][j] = scores[0][j-1] + 1
	}

	var match int
	for i := 1; i <= len(s); i++ {
		for j := 1; j <= len(t); j++ {
			if s[i-1] == t[j-1] {
				match = 0
			} else {
				match = 1
			}
			scores[i][j] = min(scores[i-1][j]+1, scores[i][j-1]+1, scores[i-1][j-1]+match)
		}
	}
	return scores[len(s)][len(t)]
}
