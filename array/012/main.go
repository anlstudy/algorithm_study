package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	res:=[]int{}
	rightSideViewHelper(root,0,res[:])
	return res
}

func rightSideViewHelper(root *TreeNode,depth int ,res []int){
	if root == nil{
		return
	}

	if len(res) == depth {
		res = append(res,root.Val)
	}

	rightSideViewHelper(root.Right,depth+1,res)
	rightSideViewHelper(root.Left,depth+1,res)
}

func main() {
	var arr = [5]int{1, 2, 3, 4, 5}
	fmt.Println(sum(arr[:]))
	fmt.Println((arr[:]))
}

func test(arr []int){
	arr = append(arr, 1)
}

func sum(arr []int) int {
	s := 0
	for i := 0; i < len(arr); i++ {
		s += arr[i]
		arr[i] =1
	}
	arr=append(arr, 2)
	return s
}