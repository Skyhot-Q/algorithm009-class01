package main

//func climbStairs(n int) int {
//	// 斐波那契数列公式
//	//(((1+sqrt5)/2)**(n+1)  + ((1-sqrt5)/2)**(n+1))/sqrt5
//
//	sqrt5 := math.Sqrt(5)
//	x := math.Pow((1+sqrt5)/2, float64(n+1))
//	y := math.Pow((1-sqrt5)/2, float64(n+1))
//	r := (x - y)/sqrt5
//	return int(math.Round(r))
//
//}

func climbStairs(n int) int {
	// 斐波那契数列公式
	//n = n-1 + n-2
	n1 := 0
	n2 := 1
	r := 0
	for i := 0; i < n; i++ {
		r = n1 + n2
		n1, n2 = n2, r
	}
	return r
}
