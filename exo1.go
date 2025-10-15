package ateliersansia

func TwoSum(nums []int, target int) (int, int) {
	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return i, j
			}
		}
	}
	return -1, -1
}