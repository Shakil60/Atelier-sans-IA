package ateliersansia

func MaxSubArray(nums []int) int {
	sommeMax := 0
	sommeAct := 0
	for i:= range nums {
		if nums[i] >= 0 {
			sommeAct = nums[i]
			for j := i + 1; j < len(nums); j++ {
				if sommeAct >= 0 {
					sommeAct += nums[j]
				}
				if sommeAct > sommeMax{
					sommeMax = sommeAct
				}
			}
		}
	}
	return sommeMax
}