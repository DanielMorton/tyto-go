package backtrack

func OutputGB(backtrack [][]string, s, t string, i, j int) (string, string) {
	if i == 0 && j == 0 {
		return "", ""
	} else if backtrack[i][j] == "U" {
		sAlign, tAlign := OutputGB(backtrack, s, t, i-1, j)
		return sAlign + string(s[i-1]), tAlign + "-"
	} else if backtrack[i][j] == "L" {
		sAlign, tAlign := OutputGB(backtrack, s, t, i, j-1)
		return sAlign + "-", tAlign + string(t[j-1])
	} else {
		sAlign, tAlign := OutputGB(backtrack, s, t, i-1, j-1)
		return sAlign + string(s[i-1]), tAlign + string(t[j-1])
	}
}

func OutputLB(backtrack [][]string, s, t string, end [2]int) (string, string) {
	if end[0] == 0 || end[1] == 0 || backtrack[end[0]][end[1]] == "S" {
		return "", ""
	} else if backtrack[end[0]][end[1]] == "U" {
		i := end[0]
		end[0]--
		sAlign, tAlign := OutputLB(backtrack, s, t, end)
		return sAlign + string(s[i-1]), tAlign + "-"
	} else if backtrack[end[0]][end[1]] == "L" {
		j := end[1]
		end[1]--
		sAlign, tAlign := OutputLB(backtrack, s, t, end)
		return sAlign + "-", tAlign + string(t[j-1])
	} else {
		i, j := end[0], end[1]
		end[0]--
		end[1]--
		sAlign, tAlign := OutputLB(backtrack, s, t, end)
		return sAlign + string(s[i-1]), tAlign + string(t[j-1])
	}
}
