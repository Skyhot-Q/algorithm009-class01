package template

type Node struct {
	val  int
	children []*Node
}

func dfs(node *Node, visited map[int]struct{}) {
	if _, ok := visited[node.val]; ok {
		return
	}
	visited[node.val] = struct{}{}

	for _, nextNode := range node.children{
		dfs(nextNode, visited)
	}
}

func bfs(graph, start, end int) {
	queue := []int{start}
	visited := make(map[int]struct{})
	for len(queue) > 0 {
		node := queue[len(queue)-1]
		visited[node] = struct{}{}
		process(node)
		nodes = generateRelatedNodes(node)
		queue.Push(nodes)
	}

}