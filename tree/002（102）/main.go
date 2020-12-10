package main

import (
	"container/list"
	"fmt"
)

//第102题：二叉树的层次遍历
//给定一个二叉树，返回其按层次遍历的节点值。（即逐层地，从左到右访问所有节点）。
//示例:
//
//给定二叉树 [3,9,20,null,null,15,7]，
//3
/// \
//9  20
///  \
//15   7
//返回其层次遍历结果：[[3],[9,20],[15,7]]

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func levelOrder (root *TreeNode) [][]int{
	var result [][]int
	if root == nil{
		return result
	}
	queue:=list.New()
	queue.PushFront(root)
	for queue.Len() > 0 {
		var cur []int
		listLength := queue.Len()
		for i := 0; i < listLength; i++ {
			node := queue.Remove(queue.Back()).(*TreeNode)
			cur = append(cur, node.Val)
			if node.Left!=nil{
				queue.PushFront(node.Left)
			}
			if node.Right!=nil{
				queue.PushFront(node.Right)
			}
		}
		result = append(result,cur)
	}
	return result
}

//func levelOrder(root *TreeNode) [][]int {
//	var result [][]int
//	if root == nil {
//		return result
//	}
//	// 定义一个双向队列
//	queue := list.New()
//	// 头部插入根节点
//	queue.PushFront(root)
//	// 进行广度搜索
//	for queue.Len() > 0 {
//		var current []int
//		listLength := queue.Len()
//		for i := 0; i < listLength; i++ {
//			// 消耗尾部
//			// queue.Remove(queue.Back()).(*TreeNode)：移除最后一个元素并将其转化为TreeNode类型
//			node := queue.Remove(queue.Back()).(*TreeNode)
//			current = append(current, node.Val)
//			if node.Left != nil {
//				//插入头部
//				queue.PushFront(node.Left)
//			}
//			if node.Right != nil {
//				queue.PushFront(node.Right)
//			}
//		}
//		result = append(result, current)
//	}
//	return result
//}

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
	fmt.Println(levelOrder(node))
}