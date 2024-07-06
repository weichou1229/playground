package golang

// removeDuplicates given an integer array nums sorted in non-decreasing order, remove the duplicates in-place such that each unique element appears only once. The relative order of the elements should be kept the same. Then return the number of unique elements in nums.
// Consider the number of unique elements of nums to be k, to get accepted, you need to do the following things:
// Change the array nums such that the first k elements of nums contain the unique elements in the order they were present in nums initially. The remaining elements of nums are not important as well as the size of nums.
// Return k.
// https://leetcode.com/problems/remove-duplicates-from-sorted-array
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return 1
	}
	var list = []int{nums[0]}
	for i := 1; i < len(nums); i++ {
		if list[len(list)-1] == nums[i] {
			continue
		}
		list = append(list, nums[i])
	}
	// replace original slice
	for i := 0; i < len(list); i++ {
		nums[i] = list[i]
	}
	nums = nums[:len(list)]
	return len(list)
}
