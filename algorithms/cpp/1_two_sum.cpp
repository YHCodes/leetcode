// Source:
// Author: lyh
// Date: 2019-11-23

/*******************************************************************
Problem Describe:

Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/two-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
********************************************************************/


/*******************************************************************
Problem Solution:
一句话描述题目: 输入: 一个不重复的数组和数字, 找出数组中的两个元素相加等于给定数字的位置

思路:由于数组是不重复的, 可以考虑用哈希。
1> 先定义一个 map, key 是数组元素的值, value保存数组元素的位置
2> 遍历数组, 每个元素的值记为 val,  看target-val是否在字典中, 如果是的话, 则返回位置, 如果不是的话, 则加入到字典中。


********************************************************************/

#include <iostream>
#include <vector>
using namespace std;

class Solution {
public:
    vector<int> twoSum(vector<int>& nums, int target) {
        map<int, int> m;
        
        for (int i = 0; i < nums.size(); ++i) {
            int dif = target - nums[i];
            if (m.find(dif) != m.end()) {
                return {m[dif], i};
            } else {
                m[nums[i]] = i;
            }
        }
        return {-1, -1};
    }
};