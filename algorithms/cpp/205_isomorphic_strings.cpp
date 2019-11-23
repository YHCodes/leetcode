// Author: lyh
// Date:

/*******************************************************************
Problem Describe:

Given two strings s and t, determine if they are isomorphic.

Two strings are isomorphic if the characters in s can be replaced to get t.

All occurrences of a character must be replaced with another character while preserving the order of characters. No two characters may map to the same character but a character may map to itself.

Example 1:

Input: s = "egg", t = "add"
Output: true
Example 2:

Input: s = "foo", t = "bar"
Output: false
Example 3:

Input: s = "paper", t = "title"
Output: true
Note:
You may assume both s and t have the same length.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/isomorphic-strings
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
********************************************************************/


/*******************************************************************
Problem Solution:
一句话描述题目: 判断两个等长的字符串是否是同构的

思路
错误的方法: 把两个字符串转换成 set, 然后判断两个 set 的大小是否相等, 然而这种方法不全面，对于 aba, baa 就无法判断
如果只用一个 hash, 那么可能有些条件忽视掉了

正确的解法1: 用两个hash

正确的解法2: 

********************************************************************/

#include <iostream>
#include <vector>
using namespace std;

class Solution {
public:


    bool isIsomorphic(string s, string t) {
        map<char, char> m1;
        map<char, char> m2;

        for (int i = 0; i < s.size(); ++i) {
            char c1 = s[i];
            char c2 = t[i];

            if (m1.find(c1) != m1.end()) {
                if (m1[c1] != c2) return false;
            } else if (m2.find(c2) != m2.end()) {
                if (m2[c2] != c1) return false;
            } else {
                m1[c1] = c2;
                m2[c2] = c1;
            }
        }
        return true;
    }
};