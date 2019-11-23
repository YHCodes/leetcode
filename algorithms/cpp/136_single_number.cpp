// Author: lyh
// Date:

/*******************************************************************
Problem Describe:

Given a non-empty array of integers, every element appears twice except for one. Find that single one.

Note:

Your algorithm should have a linear runtime complexity. Could you implement it without using extra memory?

Example 1:

Input: [2,2,1]
Output: 1
Example 2:

Input: [4,1,2,1,2]
Output: 4

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/single-number
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
********************************************************************/


/*******************************************************************
Problem Solution:
一句话描述题目: 给定一个非空数组, 求只出现一次的元素

思路:
方法1. 直接用hash, 时间复杂度 O(n), 空间复杂度O(n)
方法2. 挑战不使用额外空间
    位操作.  重复两次的元素通过位操作变成0, 异或
    C++ 异或操作符是 ^


********************************************************************/

#include <iostream>
#include <vector>
#include <map>
using namespace std;

class Solution {
public:
    // method 1, hash
    int singleNumber_1(vector<int>& nums) {
        map<int, int> m;
        int sum = 0;
        for (int i = 0; i < nums.size(); ++i) {
            if (m.find(nums[i]) != m.end()) {
                sum -= nums[i];
            } else {
                m[nums[i]] = 0;
                sum += nums[i];
            }
        }
        return sum;
    }

    // method 2, 位操作
    int singleNumber(vector<int>& nums) {
        int result = 0;

        for (int i = 0; i < nums.size(); ++i) {
            result ^= nums[i];
        }

        return result;
    }
};