package golang

// twoSun given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
// You can return the answer in any order.
func twoSum(nums []int, target int) []int {
	complementIndexMap := make(map[int]int, len(nums))
	for i, v := range nums {
		if index, ok := complementIndexMap[target-v]; ok {
			return []int{i, index}
		}
		complementIndexMap[v] = i
	}
	return []int{}
}
