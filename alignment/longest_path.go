package alignment

import "github.com/tyto-go/internal"

func LongestPathLength(m, n int, Down [][]int, Right [][]int) int {
	grid := internal.MakeGrid[int](m+1, n+1)
	var down, right int
	grid[0][0] = 0
	for i := 1; i <= m; i++ {
		grid[i][0] = grid[i-1][0] + Down[i-1][0]
	}
	for j := 1; j <= n; j++ {
		grid[0][j] = grid[0][j-1] + Right[0][j-1]
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			down = grid[i-1][j] + Down[i-1][j]
			right = grid[i][j-1] + Right[i][j-1]
			if down > right {
				grid[i][j] = down
			} else {
				grid[i][j] = right
			}
		}
	}
	return grid[m][n]
}
