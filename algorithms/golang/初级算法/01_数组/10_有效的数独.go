/*
判断一个 9x9 的数独是否有效。只需要根据以下规则，验证已经填入的数字是否有效即可。

数字 1-9 在每一行只能出现一次。
数字 1-9 在每一列只能出现一次。
数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。

示例 1:
输入:
[
  ["5","3",".",".","7",".",".",".","."],
  ["6",".",".","1","9","5",".",".","."],
  [".","9","8",".",".",".",".","6","."],
  ["8",".",".",".","6",".",".",".","3"],
  ["4",".",".","8",".","3",".",".","1"],
  ["7",".",".",".","2",".",".",".","6"],
  [".","6",".",".",".",".","2","8","."],
  [".",".",".","4","1","9",".",".","5"],
  [".",".",".",".","8",".",".","7","9"]
]
输出: true

链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2f9gg/
*/
package main

import "fmt"

func isValidSudoku(board [][]byte) bool {
    rows := [9][9]int{}
    cols := [9][9]int{}
    boxes := [9][9]int{}

    for i:= 0; i < 9; i++ {
        for j := 0; j < 9; j++ {
            if board[i][j] != '.' {
                num := board[i][j] - '1'
                box_index := (i/3) * 3 + j/3
                if rows[i][num] == 1 {
                    return false 
                } else {
                    rows[i][num] = 1
                }

                if cols[j][num] == 1 {
                    return false
                } else {
                    cols[j][num] = 1
                }

                if boxes[box_index][num] == 1 {
                    return false
                } else {
                    boxes[box_index][num] = 1
                }
            }
        }
    }
    return true
}
