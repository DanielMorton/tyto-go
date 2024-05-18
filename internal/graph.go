package internal

func MakeVertices(G [][3]int) map[int]bool {
	vertices := make(map[int]bool)
	for _, edge := range G {
		vertices[edge[0]] = true
		vertices[edge[1]] = true
	}
	return vertices
}

func MakeWeightedGraph(G [][3]int, reverse bool) map[int][]WeightedNode {
	vertices := MakeVertices(G)
	graph := make(map[int][]WeightedNode)
	for vertex := range vertices {
		graph[vertex] = make([]WeightedNode, 0)
	}
	for _, edge := range G {
		if reverse {
			graph[edge[1]] = append(graph[edge[1]], WeightedNode{Node: edge[0], Weight: edge[2]})
		} else {
			graph[edge[0]] = append(graph[edge[0]], WeightedNode{Node: edge[1], Weight: edge[2]})
		}
	}
	return graph
}
