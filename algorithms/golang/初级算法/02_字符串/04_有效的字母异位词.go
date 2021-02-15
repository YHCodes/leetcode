/*
有效的字母异位词
给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。

示例 1:
输入: s = "anagram", t = "nagaram"
输出: true

示例 2:
输入: s = "rat", t = "car"
输出: false
说明:
你可以假设字符串只包含小写字母。

进阶:
如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？

链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xn96us/
*/
package main

import "fmt"
import "sort"

// 方法一：排序, 时间复杂度O(nlogn) 空间复杂度O(logn) 
func isAnagram(s string, t string) bool {
	s1 := []byte(s)
	s2 := []byte(t)
	sort.Slice(s1, func(i, j int) bool {
		return s1[i] < s1[j]
	})
	sort.Slice(s2, func(i, j int) bool {
		return s2[i] < s2[j]
	})
	return string(s1) == string(s2)
}

// 方法二：hashTable 官方解答 (推荐)
func isAnagram2(s string, t string) bool {
	var c1, c2 [26]int
	for _, ch := range s {
		c1[ch - 'a']++
	}
	for _, ch := range t {
		c2[ch - 'a']++
	}
	return c1 == c2
}

// 方法二: 我的解法, hashTable  时间复杂度O(n)  空间复杂度O(n)
func isAnagram3(s string, t string) bool {
	// build hash table1
	m1 := map[byte]int{}
	for idx := 0; idx < len(s); idx++ {
		m1[s[idx]]++
	}
	// build hash table2
	m2 := map[byte]int{}
	for idx := 0; idx < len(t); idx++ {
		m2[t[idx]]++
	}
	// compare
	if len(m1) != len(m2) {
		return false
	}
	for ch, cnt := range m1 {
		if cnt != m2[ch] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isAnagram2("anagram", "nagaram"))
	fmt.Println(isAnagram2("rat", "car"))
}