package twosum

func twoSum(nums []int, target int) []int {
	arr := []int{0, 0}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				arr[0] = i
				arr[1] = j
				return arr
			}
		}
	}
	return arr
}
