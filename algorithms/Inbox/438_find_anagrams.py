#encoding="utf8"
from typing import List
"""
给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。

异位词：相同字母重排列形成的字符串

示例 1:
输入: s = "cbaebabacd", p = "abc"
输出: [0,6]
解释:
起始索引等于 0 的子串是 "cba", 它是 "abc" 的异位词。
起始索引等于 6 的子串是 "bac", 它是 "abc" 的异位词。

示例 2:
输入: s = "abab", p = "ab"
输出: [0,1,2]
解释:
起始索引等于 0 的子串是 "ab", 它是 "ab" 的异位词。
起始索引等于 1 的子串是 "ba", 它是 "ab" 的异位词。
起始索引等于 2 的子串是 "ab", 它是 "ab" 的异位词。

提示:
1 <= s.length, p.length <= 3 * 10^4
s 和 p 仅包含小写字母
"""

class Solution:
    def findAnagrams(self, s: str, p: str) -> List[int]:
        ans = []
        if len(s) < len(p):
            return ans
        
        s_counter = [0] * 26
        p_counter = [0] * 26

        for i in range(len(p)):
            s_counter[ord(s[i]) - 97] += 1
            p_counter[ord(p[i]) - 97] += 1
        
        if s_counter == p_counter:
            ans.append(0)
        
        for i in range(len(s) - len(p)):
            s_counter[ord(s[i]) - 97] -= 1
            s_counter[ord(s[len(p) + i]) - 97] += 1

            if s_counter == p_counter:
                ans.append(i + 1)
        return ans
        
if __name__ == "__main__":
    s1 = "cbaebabacd"
    p1 = "abc"
    assert Solution().findAnagrams(s1, p1) == [0, 6]

    s2 = "abab"
    p2 = "ab"
    assert Solution().findAnagrams(s1, p1) == [0, 1, 2]