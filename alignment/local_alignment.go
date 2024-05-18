package alignment

import "github.com/tyto-go/alignment/backtrack"

func LocalAlignment(match int, mismatch int, gap int, s string, t string) (int, string, string) {
	score, end, bt := backtrack.LocalBacktrack(match, mismatch, gap, s, t)
	sAlign, tAlign := backtrack.OutputLB(bt, s, t, end)
	return score, sAlign, tAlign
}
