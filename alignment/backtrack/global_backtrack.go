package backtrack

import "github.com/tyto-go/internal"

func GlobalBacktrack(match, mismatch, gap int, s, t string) (int, [][]string) {
	scores := internal.MakeGrid[int](len(s)+1, len(t)+1)
	backtrack := internal.MakeGrid[string](len(s)+1, len(t)+1)

	scores[0][0] = 0
	for i := 1; i <= len(s); i++ {
		scores[i][0] = scores[i-1][0] - gap
		backtrack[i][0] = "U"
	}
	for j := 1; j <= len(t); j++ {
		scores[0][j] = scores[0][j-1] - gap
		backtrack[0][j] = "L"
	}

	var matchScore int
	for i := 1; i <= len(s); i++ {
		for j := 1; j <= len(t); j++ {
			if s[i-1] == t[j-1] {
				matchScore = match
			} else {
				matchScore = -mismatch
			}
			scores[i][j] = max(scores[i-1][j]-gap, scores[i][j-1]-gap,
				scores[i-1][j-1]+matchScore)
			if scores[i][j] == scores[i-1][j]-gap {
				backtrack[i][j] = "U"
			} else if scores[i][j] == scores[i][j-1]-gap {
				backtrack[i][j] = "L"
			} else {
				backtrack[i][j] = "D"
			}
		}
	}
	return scores[len(s)][len(t)], backtrack
}
