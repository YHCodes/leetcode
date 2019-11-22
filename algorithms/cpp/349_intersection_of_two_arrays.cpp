// Author: lyh
// Date: 2019-11-22

/*******************************************************************
Problem Describe:

Given two arrays, write a function to compute their intersection.

Example 1:
Input: nums1 = [1,2,2,1], nums2 = [2,2]
Output: [2]

Example 2:
Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
Output: [9,4]

Note:
Each element in the result must be unique.
The result can be in any order.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/intersection-of-two-arrays
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
********************************************************************/


/*******************************************************************
Problem Solution:

给予两个数组, 求它们的交集
- 思路1, 暴力法, 用两个循环逐一对比 O(n2)
- 思路2, 先对两个数组排序, 然后对比
- 思路3, 利用 set
********************************************************************/

#include <iostream>
#include <vector>
#include <set>
using namespace std;

class Solution {
public:
    vector<int> intersection(vector<int>& nums1, vector<int>& nums2) {
        vector<int> result;
        set<int> s1;
        set<int> s2;

        for (int i = 0; i < nums1.size(); ++i) {
            s1.insert(nums1[i]);
        }

        for (int i = 0; i < nums2.size(); ++i) {
            s2.insert(nums2[i]);
        }

        set<int>::iterator it;
        for (it = s1.begin(); it != s1.end(); it++) {
            if (s2.find(*it) != s2.end()) {
                result.push_back(*it);
            }
        }

        return result;
    }
};