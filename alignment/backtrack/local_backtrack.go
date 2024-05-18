package backtrack

import "github.com/tyto-go/internal"

func LocalBacktrack(match, mismatch, gap int, s, t string) (int, [2]int, [][]string) {
	scores := internal.MakeGrid[int](len(s)+1, len(t)+1)
	backtrack := internal.MakeGrid[string](len(s)+1, len(t)+1)

	scores[0][0] = 0
	for i := 1; i <= len(s); i++ {
		scores[i][0] = 0
		backtrack[i][0] = "S"
	}
	for j := 1; j <= len(t); j++ {
		scores[0][j] = 0
		backtrack[0][j] = "S"
	}

	var matchScore int
	var end [2]int
	maxScore := 0
	for i := 1; i <= len(s); i++ {
		for j := 1; j <= len(t); j++ {
			if s[i-1] == t[j-1] {
				matchScore = match
			} else {
				matchScore = -mismatch
			}
			scores[i][j] = max(scores[i-1][j]-gap, scores[i][j-1]-gap,
				scores[i-1][j-1]+matchScore, 0)
			if scores[i][j] > maxScore {
				maxScore = scores[i][j]
				end = [2]int{i, j}
			}
			if scores[i][j] == 0 {
				backtrack[i][j] = "S"
			} else if scores[i][j] == scores[i-1][j]-gap {
				backtrack[i][j] = "U"
			} else if scores[i][j] == scores[i][j-1]-gap {
				backtrack[i][j] = "L"
			} else {
				backtrack[i][j] = "D"
			}
		}
	}
	return maxScore, end, backtrack
}
