/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除有序数组中的重复项
 */
package leetcode

// @lc code=start
func removeDuplicates(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return 1
	}
	p, q := 0, 1
	for ; q < len(nums); q++ {
		if nums[p] != nums[q] {
			nums[p+1] = nums[q]
			p++
		}
	}
	return p + 1
}

// @lc code=end
