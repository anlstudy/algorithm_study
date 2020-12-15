package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var maxres int

func maxPathSum(root *TreeNode) int {
	maxres = math.MinInt64
	maxPathSumHelp(root)
	return maxres
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxPathSumHelp(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := maxInt(maxPathSumHelp(root.Left), 0)
	right := maxInt(maxPathSumHelp(root.Right), 0)
	maxres = maxInt(maxres, root.Val+left+right)
	return maxInt(left,right)+root.Val
}

func main() {
	node:=&TreeNode{
		Val: -2,
		Left: &TreeNode{
			Val: 1,
		},
	}
	fmt.Println(maxPathSum(node))
}
