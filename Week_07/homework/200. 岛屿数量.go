package main

type UnionFind struct {
	count  int
	parent []int
}

func NewUnionFind(grid [][]byte) *UnionFind {
	uf := new(UnionFind)
	n := len(grid)
	m := len(grid[0])
	parent := make([]int, m*n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				parent[i*n+j] = i*n + j
				uf.count++
			}
		}
	}
	uf.parent = parent
	return uf
}

func (uf *UnionFind) union(e1, e2 int) {
	p1 := uf.find(e1)
	p2 := uf.find(e2)
	if p1 != p2 {
		uf.parent[p1] = p2
		uf.count--
	}
}

func (uf *UnionFind) find(input int) int {
	if uf.parent[input] != input {
		uf.parent[input] = uf.find(uf.parent[input])
	}
	return uf.parent[input]
}

func (uf *UnionFind) getCount() int {
	return uf.count
}

func numIslands(grid [][]byte) int {
	//1.如何建立这个矩阵
	if grid == nil || len(grid) == 0 {
		return 0
	}
	row := len(grid)
	col := len(grid[0])
	uf := NewUnionFind(grid)
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == '1' {
				grid[i][j] = '0'
				cur := i*col + j
				if i+1 < row && grid[i+1][j] == '1' {
					uf.union(cur, (i+1)*col+j)
				}

				if j+1 < col && grid[i][j+1] == '1' {
					uf.union(cur, i*col+j+1)
				}
			}
		}
	}
	return uf.count
}
