/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */
package leetcode

// @lc code=start
func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for i, v := range nums {
		if j, ok := m[target-v]; ok {
			return []int{i, j}
		}
		m[v] = i
	}
	return nil
}

// @lc code=end
