package main

import "fmt"

//第104题：二叉树的最大深度
//给定一个二叉树，找出其最大深度。二叉树的深度为根节点到最远叶子节点的最长路径上的节点数
//说明: 叶子节点是指没有子节点的节点。
//
//示例：
//
//给定二叉树 [3,9,20,null,null,15,7]，
//3
/// \
//9  20
///  \
//15   7

type TreeNode struct {
     Val int
     Left *TreeNode
	 Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right))+1
}

func max(a int,b int) int {
	if a > b {
		return a
	}
	return b
}

func main(){
	node:=&TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
			Left: &TreeNode{
				Val: 15,
				Left: nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val: 7,
				Left: nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val: 20,
			Left: nil,
			Right: nil,
		},
	}

	fmt.Println(maxDepth(node))
}