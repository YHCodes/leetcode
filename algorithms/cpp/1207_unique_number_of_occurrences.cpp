// Author: lyh
// Date:

/*******************************************************************
Problem Describe:
Given an array of integers arr, write a function that returns true if and 
only if the number of occurrences of each value in the array is unique.

Example 1:

Input: arr = [1,2,2,1,1,3]
Output: true
Explanation: The value 1 has 3 occurrences, 2 has 2 and 3 has 1. 
No two values have the same number of occurrences.
Example 2:

Input: arr = [1,2]
Output: false
Example 3:

Input: arr = [-3,0,1,-3,1,1,1,-3,10,0]
Output: true
 

Constraints:

1 <= arr.length <= 1000
-1000 <= arr[i] <= 1000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/unique-number-of-occurrences
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。


********************************************************************/


/*******************************************************************
Problem Solution:
一句话描述题目: 数组中的元素出现次数都是不同的


********************************************************************/

#include <iostream>
#include <vector>
using namespace std;

class Solution {
public:
    bool uniqueOccurrences(vector<int>& arr) {
        map<int, int> m;
        for (int i = 0; i < arr.size(); ++i) {
            if (m.find(arr[i]) != m.end()) {
                m[arr[i]] += 1;
            } else {
                m[arr[i]] = 1;
            }
        }

        map<int, int>::iterator it = m.begin();

        set<int> s;
        while(it != m.end()) {
            s.insert(it->second);
            it++;
        }

        return m.size() == s.size();
    }
};