/*
给定一个整数数组，判断是否存在重复元素。
如果存在一值在数组中出现至少两次，函数返回 true 。如果数组中每个元素都不相同，则返回 false 。

示例 1:
输入: [1,2,3,1]
输出: true

链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x248f5/
*/

package main

import "fmt"
import "sort"

// 方法一: 排序法, 时间O(nlogn)
func containsDuplicate(nums []int) bool {
	sort.Ints(nums)
	for idx := 1; idx < len(nums); idx++ {
		if nums[idx] == nums[idx-1] {
			return true
		}
	}
	return false
}

// 方法二: HashSet
func containsDuplicate2(nums []int) bool {
	/*
	m := make(map[int]int)
	for _, val := range nums {
		_, ok := m[val]
		if ok {
			m[val] += 1
		} else {
			m[val] = 1
		}
	}
	return len(m) != len(nums)
	*/
	set := map[int]struct{}{}
	for _, v := range nums {
		if _, has := set[v]; has {
			return true
		} 
		set[v] = struct{}{}
	}
	return false
}

func main() {
	nums1 := []int{1,2,3,1}
	fmt.Println(containsDuplicate2(nums1))
}