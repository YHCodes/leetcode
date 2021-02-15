/*
给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。

进阶：
尽可能想出更多的解决方案，至少有三种不同的方法可以解决这个问题。
你可以使用空间复杂度为 O(1) 的 原地 算法解决这个问题吗？
 
示例 1:
输入: nums = [1,2,3,4,5,6,7], k = 3
输出: [5,6,7,1,2,3,4]
解释:
向右旋转 1 步: [7,1,2,3,4,5,6]
向右旋转 2 步: [6,7,1,2,3,4,5]
向右旋转 3 步: [5,6,7,1,2,3,4]

链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2skh7/
*/
package main

import "fmt"

// 方法一: 使用额外的数组, 遍历原数组, 将原数组下标为i的元素放至新数组下标为 (i+k) mod n
// 时间复杂度 O(n), 空间复杂度 O(n)
func rotate(nums []int, k int)  {
	n := len(nums)
	newNums := make([]int, n)
	for i, v := range nums {
		newNums[(i+k)%n] = v
	}
	copy(nums, newNums)
}

// 方法二：环状替换 https://leetcode-cn.com/problems/rotate-array/solution/xuan-zhuan-shu-zu-by-leetcode-solution-nipk/
// func rotate2(nums []int, k int) {
// }

// 方法三: 数组翻转, 时间O(n) 空间O(1)
// 翻转整个: [1,2,3,4,5,6,7] -> [7,6,5,4,3,2,1]
// 翻转左边k个: [5,6,7,4,3,2,1]
// 翻转右边: [5,6,7,1,2,3,4]
func reverse(nums []int) {
	for i, n := 0, len(nums); i < n/2; i++ {
		nums[i], nums[n-1-i] = nums[n-1-i], nums[i]
	}
}

func rotate3(nums []int, k int) {
	k %= len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

func main() {
	nums1 := []int{1,2,3,4,5,6,7}
	rotate3(nums1, 3)
	fmt.Println(nums1)
}