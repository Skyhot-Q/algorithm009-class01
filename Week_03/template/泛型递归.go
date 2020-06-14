package template

func recursion(level, param1, param2 interface{}) {
	// T: 终止条件
	if level > MaxLevel{
		print_process_result
		return
	}
	// P：处理当前层数据逻辑
	process(level, data)

	// D：进行下钻递归操作
	recursion(level+1, p1, ...)

	// R：恢复本层状态
}
