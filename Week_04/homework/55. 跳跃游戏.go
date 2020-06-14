package main

func canJump(nums []int) bool {
	numLen := len(nums)
	if numLen == 0 {
		return false
	}

	end := numLen - 1
	for i:=numLen-1; i >=0; i--  {
		if nums[i] +i >=end{
			end = i
		}
	}
	return end == 0
}
