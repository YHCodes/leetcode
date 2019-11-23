// Source:
// Author: lyh
// Date: 2019-11-23

/*******************************************************************
Problem Describe:
Given an array of size n, find the majority element. The majority element is the element that appears more than ⌊ n/2 ⌋ times.

You may assume that the array is non-empty and the majority element always exist in the array.

Example 1:

Input: [3,2,3]
Output: 3
Example 2:

Input: [2,2,1,1,1,2,2]
Output: 2

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/majority-element
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
********************************************************************/


/*******************************************************************
Problem Solution:

题目: 求众数

思路1. 哈希
1> 第一次遍历, 转哈希表, key是数组的元素, value 是次数
2> 第二次遍历, 求众数
复杂度: 时间复杂度O(n) 空间复杂度O(n)

思路2. 排序
1> 排序
2> 返回 nums[nums.size()/2]
复杂度: 时间复杂度 O(nlogn), 空间复杂度O(1)

********************************************************************/

#include <iostream>
#include <vector>
#include <map>
using namespace std;

class Solution {
public:
    int majorityElement(vector<int>& nums) {
        map<int, int> map_nums;

        for (int i = 0; i < nums.size(); ++i) {
            if (map_nums.find(nums[i]) == map_nums.end()) {
                map_nums[nums[i]] = 1;
            } else {
                map_nums[nums[i]] += 1;
            }
        }

        map<int, int>::iterator it;
        int thresh = nums.size() / 2;
        for (it = map_nums.begin(); it != map_nums.end(); ++it) {
            if (it->second > thresh) {
                return it->first;
            }
        }
        return -1;
    }
};