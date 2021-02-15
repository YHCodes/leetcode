/*
给定两个数组，编写一个函数来计算它们的交集。

示例 1：
输入：nums1 = [1,2,2,1], nums2 = [2,2]
输出：[2,2]

示例 2:
输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
输出：[4,9]

链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2y0c2/
*/
package main

import "fmt"
import "sort"

// HashMap优化版: 构建哈希表时, 考虑两个输入数组的大小, 减少空间消耗
func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		return intersect(nums2, nums1)
	}

	m := map[int]int{}
	for _, num := range nums1 {
		m[num]++
	}
	
	intersection := []int{}
	for _, num := range nums2 {
		if m[num] > 0 {
			intersection = append(intersection, num)
			m[num]--
		}
	}
	return intersection
}

// 法2: 排序法
func intersect2(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	
	result := []int{}
	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		if nums1[i] == nums2[j] {
			result = append(result, nums1[i])
			i += 1
			j += 1
		} else if nums1[i] < nums2[j] {
			i += 1
		} else {
			j += 1
		}
	}
	return result
}

// 其它: HashMap 普通版
func intersect_bad(nums1 []int, nums2 []int) []int {
	m := map[int]int{}
	for _, val := range nums1 {
		if _, has := m[val]; has {
			m[val] += 1
		} else {
			m[val] = 1
		}
	}

	result := []int{}
	for _, val := range nums2 {
		if cnt, has := m[val]; has && cnt > 0 {
			m[val] -= 1
			result = append(result, val)
		}
	}
	return result
}

func main() {
	nums1 := []int{1,2,2,1}
	nums2 := []int{2,2}
	fmt.Println(intersect(nums1, nums2))

	arr1 := []int{4,9,5}
	arr2 := []int{9,4,9,8,4}
	fmt.Println(intersect(arr1, arr2))
}