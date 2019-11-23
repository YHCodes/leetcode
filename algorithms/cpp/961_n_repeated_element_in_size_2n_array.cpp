
// Author: lyh
// Date: 2019-11-23

/*******************************************************************
Problem Describe:
In a array A of size 2N, there are N+1 unique elements, 
and exactly one of these elements is repeated N times.

Return the element repeated N times.

Example 1:
Input: [1,2,3,3]
Output: 3

Example 2:
Input: [2,1,2,5,3,2]
Output: 2

Example 3:
Input: [5,1,5,2,5,3,5,4]
Output: 5
 

Note:

4 <= A.length <= 10000
0 <= A[i] < 10000
A.length is even

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/n-repeated-element-in-size-2n-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
********************************************************************/


/*******************************************************************
Problem Solution:
一句话描述题目: 找出2N大小数组中, 其中出现N次的元素


********************************************************************/
class Solution {
public:
    int repeatedNTimes(vector<int>& A) {
        int len = A.size();
        int n = len / 2;

        map<int, int> m;
        for (int i = 0; i < A.size(); ++i) {
            if (m.find(A[i]) != m.end()) {
                m[A[i]] +=1;
            } else {
                m[A[i]] = 1;
            }
        }


        map<int, int>::iterator it = m.begin();
        while (it != m.end()) {
            if (it->second == n) {
                return it->first;
            }
            it++;
        }
        return -1;
    }
};