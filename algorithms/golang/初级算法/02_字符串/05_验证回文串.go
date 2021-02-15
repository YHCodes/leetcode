/*
验证回文串
给定一个字符串，验证它是否是回文串，只考虑字母和数字符，可以忽略字母的大小写。
说明：本题中，我们将空字符串定义为有效的回文串。

示例 1:
输入: "A man, a plan, a canal: Panama"
输出: true

示例 2:
输入: "race a car"
输出: false

链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xne8id/
*/

package main

import "fmt"
import "strings"

// 双指针法
func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	left, right := 0, len(s) - 1
	for left < right {
		if !isalnum(s[left]) {
			left += 1
			continue
		}

		if !isalnum(s[right]) {
			right -= 1
			continue
		}

		if s[left] != s[right] {
			return false
		}

		left += 1
		right -= 1
	}
	return true
}

func isalnum(ch byte) bool {
	return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome("race a car"))
	fmt.Println(isPalindrome("0P"))
}