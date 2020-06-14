package template

func dfs(node int, visited map[int]struct{}){
	if _, ok := visited[node]; ok {
		return
	}

	visited[node]= struct{}{}

	for _, next_node := range node.children(){
		if _, ok := visited[next_node]; ok {
			dfs(next_node, visited)
		}
	}
}