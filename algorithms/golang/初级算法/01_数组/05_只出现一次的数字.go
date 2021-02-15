/*
给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
说明：
你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？

示例 1:
输入: [2,2,1]
输出: 1

示例 2:
输入: [4,1,2,1,2]
输出: 4

链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x21ib6/
*/
package main

import "fmt"

// 法一: 异或运算.. 时间O(n), 空间O(1)
func singleNumber(nums []int) int {
	cur := nums[0]
	for _, val := range nums[1:] {
		cur ^= val
	}
	return cur
}

// 法二: HashMa: 时间O(n) 空间O(n)
func singleNumber2(nums []int) int {
	m := map[int]int{}
	for _, val := range nums {
		if _, has := m[val]; has {
			m[val] += 1
		} else {
			m[val] = 1
		}
	}
	
	for k, v := range m {
		if v == 1 {
			return k
		}
	}
	return -1
}

func main() {
	nums1 := []int{2,2,1}
	fmt.Println(singleNumber(nums1))
}