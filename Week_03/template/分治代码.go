package template

func divide_conquer(problem, param1, param2 interface{}, ...) {
	// T: 终止条件,问题解决了
	if problem == None {
		print_result
		return
	}

	// P：处理当前层数据逻辑,拆分子问题，怎么拆解子问题很重要
	data := prepare_data(problem)
	suproblems := split_problem(problem, data)

	// D：进行下钻递归操作
	subresult1 := divide_conquer(suproblem[0], p1, ...)
	subresult2 := divide_conquer(suproblem[1], p1, ...)
	subresult3 := divide_conquer(suproblem[2], p1, ...)
	...

	// M: 整合结果数据，怎么汇聚结果
	result := process_result(subresult1, subresult2, subresult3, ...)

	// R：恢复本层状态
}
