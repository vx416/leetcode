package array

// https://leetcode.com/problems/find-the-duplicate-number/
func findDuplicate(nums []int) int {
	slow := 0
	fast := 0

	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}

	slow = 0

	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}

	return slow
}
