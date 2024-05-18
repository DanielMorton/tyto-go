package alignment

import "github.com/tyto-go/alignment/backtrack"

func GlobalAlignment(match int, mismatch int, gap int, s string, t string) (int, string, string) {
	score, bt := backtrack.GlobalBacktrack(match, mismatch, gap, s, t)
	sAlign, tAlign := backtrack.OutputGB(bt, s, t, len(s), len(t))
	return score, sAlign, tAlign
}
