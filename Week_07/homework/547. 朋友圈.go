package main

type UnionFind struct {
	count  int
	parent []int
}

func NewUnionFind(n int) *UnionFind{
	u := UnionFind{
		count:  n,
		parent: make([]int, n),
	}
	for i := 0; i < n; i++ {
		u.parent[i] = i
	}
	return &u
}

func (u *UnionFind) find(p int) int{
	for p != u.parent[p] {
		u.parent[p]  = u.parent[u.parent[p]]
		p = u.parent[p]
	}
	return p
}

func (u *UnionFind) union(p, q int){
	rootP := u.find(p)
	rootQ := u.find(q)
	if rootP == rootQ {
		return
	}
	u.parent[rootP] = rootQ
	u.count--
}


func findCircleNum(M [][]int) int {
	if M == nil {
		return 0
	}
	n := len(M)
	u := NewUnionFind(n)
	for i:=0; i<n;i++ {
		for j:=0; j< n;j++ {
			if M[i][j] == 1 {
				 u.union(i, j)
			}
		}
	}
	ans := make(map[int]struct{}, 0)
	for k :=0; k < n ;k++ {
		t := u.find(k)
		ans[t] = struct{}{}
	}
	return len(ans)
}
