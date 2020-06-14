package main

import "fmt"

func dfs(grid *[][]byte, i, j int) {
	nr := len(*grid)
	tmp := *grid
	nc := len(tmp[0])
	if i < 0 || j < 0 || i >= nr || j >= nc || tmp[i][j] == '0' {
		return
	}
	tmp[i][j] = '0'
	grid = &tmp
	dfs(grid, i+1, j)
	dfs(grid, i-1, j)
	dfs(grid, i, j+1)
	dfs(grid, i, j-1)
}

func numIslands(grid [][]byte) int {
	// 遇到1时就是有岛屿
	// 然后将1当做原点向外铺开，直至各方为0
	// 循环遍历数组重复操作
	nrow := len(grid)
	if nrow == 0 {
		return 0
	}
	ncol := len(grid[0])
	islandNum := 0
	for i := 0; i < nrow; i++ {
		for j := 0; j < ncol; j++ {
			if grid[i][j] == '1' {
				dfs(&grid, i, j)
				islandNum++
			}
		}
	}
	return islandNum
}

func main() {
	testCases := [][][]byte{
		{
			{'1', '1', '1', '1', '0'},
			{'1', '1', '0', '1', '0'},
			{'1', '1', '0', '0', '0'},
			{'0', '0', '0', '0', '0'},
		},
		{
			{'1', '1', '0', '0', '0'},
			{'1', '1', '0', '0', '0'},
			{'0', '0', '1', '0', '0'},
			{'0', '0', '0', '1', '1'},
		}}
	for _, c := range testCases {
		r := numIslands(c)
		fmt.Println("Result:", r)
	}
}
