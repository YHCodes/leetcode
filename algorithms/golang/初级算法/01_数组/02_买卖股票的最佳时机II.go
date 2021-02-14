package main

import "fmt"

/*
给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。
设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。
注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。


示例 1:
输入: [7,1,5,3,6,4]
输出: 7
解释: 在第 2 天（股票价格 = 1）的时候买入，在第 3 天（股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。
     随后，在第 4 天（股票价格 = 3）的时候买入，在第 5 天（股票价格 = 6）的时候卖出, 这笔交易所能获得利润 = 6-3 = 3 。

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2zsx1/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

func maxProfit(prices []int) int {
	result := 0
	idx := 1
	for ; idx < len(prices); idx++ {
		if prices[idx] > prices[idx - 1] {
			result += (prices[idx] - prices[idx - 1])
		}
	}
	return result
}

func main() {
	prices1 := []int{7,1,5,3,6,4}
	fmt.Println(maxProfit(prices1) == 7)

	prices2 := []int{1,2,3,4,5}
	fmt.Println(maxProfit(prices2) == 4)

	prices3 := []int{7,6,4,3,1}
	fmt.Println(maxProfit(prices3) == 0)
}