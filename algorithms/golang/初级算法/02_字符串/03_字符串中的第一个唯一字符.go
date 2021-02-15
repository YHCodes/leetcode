/*
字符串中的第一个唯一字符
给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。


示例：
s = "leetcode"
返回 0

s = "loveleetcode"
返回 2
 
提示：你可以假定该字符串只包含小写字母。

链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xn5z8r/
*/
package main

import "fmt"

// 方法：hash
func firstUniqChar(s string) int {
	hashTable := [26]int{}
	for _, ch := range(s) {
		hashTable[ch - 'a']++
	}
	
	for idx, ch := range(s) {
		if hashTable[ch - 'a'] == 1 {
			return idx
		}
	}
	return -1
}

func main() {
	fmt.Println(firstUniqChar("leetcode") == 0)
	fmt.Println(firstUniqChar("loveleetcode") == 2)
}