/*
谷歌：我们90％的工程师使用您编写的软件(Homebrew)，
但是您却无法在面试时在白板上写出翻转二叉树这道题，这太糟糕了。

翻转一棵二叉树。

示例：
输入：
     4
   /   \
  2     7
 / \   / \
1   3 6   9

输出：
     4
   /   \
  7     2
 / \   / \
9   6 3   1

链接：https://leetcode-cn.com/problems/invert-binary-tree
*/

package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

/*
	递归法:
	从根节点开始, 递归地对树进行遍历, 并从叶子结点先开始翻转。
	如果当前遍历到的节点root的左右两棵子树都已经翻转，那么我们只需要交换两棵子树的位置。

	python 解法:
	def invertTree(self, root: TreeNode) -> TreeNode:
		if not root:
			return root
		left = self.invertTree(root.left)
		right = self.invertTree(root.right)
		root.left, root.right = right, left
		return root
*/
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := invertTree(root.Left)
	right := invertTree(root.Right)
	root.Left = right
	root.Right = left
	return root
}