package distance

func HammingDistance(s string, t string) int {
	dist := 0
	for i := 0; i < len(s); i++ {
		if s[i] != t[i] {
			dist++
		}
	}
	return dist
}
