package main

import "fmt"

type UnionFind struct {
	count  int
	parent []int
}

func NewUnionFind(board [][]byte) *UnionFind {
	uf := new(UnionFind)
	n := len(board)
	m := len(board[0])
	parent := make([]int, m*n)
	for i :=0; i < n*m ; i ++ {
		parent[i] = i

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

func (uf *UnionFind) find(p int) int {
	if uf.parent[p] != p {
		uf.parent[p] = uf.find(uf.parent[p])
	}
	return uf.parent[p]
}

func solve(board [][]byte) {
	if board == nil || len(board) == 0 {
		return
	}
	rows := len(board)
	cols := len(board[0])
	uf := NewUnionFind(board)
	dummyNode := rows * cols
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] == 'O' {
				if i == 0 || i == rows-1 || j == 0 || j == cols-1 {
					uf.union(i*cols+j, dummyNode-1)
				} else {
					// 和上下左右合并成一个连通区域.
					//board[i][j] = 'X'
					cur := i*cols + j
					if i > 0 && board[i-1][j] == 'O' {
						uf.union(cur, (i-1)*cols+j)
					}
					if i < rows-1 && board[i+1][j] == 'O' {
						uf.union(cur, (i+1)*cols+j)
					}
					if j > 0 && board[i][j-1] == 'O' {
						uf.union(cur, i*cols+j-1)
					}
					if j < cols-1 && board[i][j+1] == 'O' {
						uf.union(cur, i*cols+j+1)
					}
				}
			}
		}
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if uf.find(i*cols+j) != uf.find(dummyNode-1) {
				board[i][j] = 'X'
			}
		}
	}
}


func main() {
	//data := [][]byte{{'X','X','X','X'},{'X','O','O','X'},{'X','X','O','X'},{'X','O','X','X'}}
	data := [][]byte{{'X','O','X','O','X','O'},{'O','X','O','X','O','X'},{'X','O','X','O','X','O'},{'O','X','O','X','O','X'}}
	solve(data)
	fmt.Println(data)
}

//["X","O","X","O","X","O"],
//["O","X","O","X","O","X"],
//["X","O","X","O","X","O"],
//["O","X","O","X","O","X"]


//["X","O","X","O","X","O"],
//["O","X","X","X","X","X"],
//["X","X","X","X","X","O"],
//["O","X","O","X","O","X"]
