package template

func DFS(tree treeNode) {
	if tree.Root == nil {
		return
	}
	visited, stack = map[int]struct{}{}, []*treeNode{tree.Root}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		visited[node] = struct{}{}

		process(node)
		nodes = generate_related_nodes(node)
		stack = append(stack, nodes...)
	}

}
