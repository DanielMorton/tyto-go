package alignment

import "github.com/tyto-go/internal"

func MakePath(prevNode map[int]int, s, t int) []int {
	path := make([]int, 0)
	node := t
	path = append(path, node)
	for node != s {
		node = prevNode[node]
		path = append(path, node)
	}
	internal.Reverse(path)
	return path
}

func LongestPath(G [][3]int, s int, t int) (int, []int) {
	graph := internal.MakeWeightedGraph(G, false)
	reverseGraph := internal.MakeWeightedGraph(G, true)

	var node, gw int
	graphWeight := make(map[int]int)
	prevNode := make(map[int]int)
	queue := make([]int, 1)
	queue[0] = s
	visited := make(map[int]bool)
	for len(queue) > 0 {
		node = queue[0]
		visited[node] = true
		queue = queue[1:]
		gw = 0
		for _, p := range reverseGraph[node] {
			if visited[p.Node] && graphWeight[p.Node]+p.Weight > gw {
				gw = graphWeight[p.Node] + p.Weight
				prevNode[node] = p.Node
			}
		}
		graphWeight[node] = gw
		for _, p := range graph[node] {
			queue = append(queue, p.Node)
		}
	}

	length := graphWeight[t]
	path := MakePath(prevNode, s, t)
	return length, path
}
