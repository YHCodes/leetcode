/*
最长公共前缀
编写一个函数来查找字符串数组中的最长公共前缀。
如果不存在公共前缀，返回空字符串 ""。

示例 1：
输入：strs = ["flower","flow","flight"]
输出："fl"

示例 2：
输入：strs = ["dog","racecar","car"]
输出：""
解释：输入不存在公共前缀。

提示：
0 <= strs.length <= 200
0 <= strs[i].length <= 200
strs[i] 仅由小写英文字母组成

链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnmav1/
*/
package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	// 可去掉
	// if len(strs) == 1 {
	// 	return strs[0]
	// }

	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			// if len(strs[j]) == 0 || len(strs[j]) - 1 < i || strs[j][i] != strs[0][i] {
			if i == len(strs[j]) || strs[j][i] != strs[0][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}

func main() {
	strs := []string{"flower","flow","flight"}
	fmt.Println(longestCommonPrefix(strs))

	strs2 := []string{"dog","racecar","car"}
	fmt.Println(longestCommonPrefix(strs2))

	strs3 := []string{"","b"}
	fmt.Println(longestCommonPrefix(strs3))

	strs4 := []string{"flower","flower","flower","flower"}
	fmt.Println(longestCommonPrefix(strs4))
}