package main

import "fmt"

func robotSim(commands []int, obstacles [][]int) int {
	dx := [4]int{0, 1, 0, -1}
	dy := [4]int{1, 0, -1, 0}
	x, y, di := 0, 0, 0

	obstacleSet := make(map[[2]int]struct{})

	for _, o := range obstacles {
		obstacleSet[[2]int{o[0], o[1]}] = struct{}{}
	}

	var ans int
	for _, c := range commands {
		if c == -2 {
			di = (di + 3) % 4
		} else if c == -1 {
			di = (di + 1) % 4
		} else {
			for k := 0; k < c; k++ {
				nx := x + dx[di]
				ny := y + dy[di]
				if _, ok := obstacleSet[[2]int{nx, ny}]; !ok {
					x = nx
					y = ny
					t := x*x + y*y
					if t > ans {
						ans = t
					}
				}
			}
		}
	}
	return ans
}

func main() {
	testCases := [][][][]int{
		//{{{4, -1, 3}}, {}},
		//{{{4, -1, 4, -2, 4}}, {{2, 4}}},
		{{{-2, 8, 3, 7, -1}}, {{-4, -1}, {1, -1}, {1, 4}, {5, 0}, {4, 5}, {-2, -1}, {2, -5}, {5, 1}, {-3, -1}, {5, -3}}},
	}
	for _, c := range testCases {
		r := robotSim(c[0][0], c[1])
		fmt.Println("Result:", r)
	}

}
