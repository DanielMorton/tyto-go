package internal

func MakeGrid[K any](x, y int) [][]K {
	grid := make([][]K, x)
	for i := range grid {
		grid[i] = make([]K, y)
	}
	return grid
}
