package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}


// 单链表快速排序 Time Complexity: O(nlogn), Space: O(n)  保证每一次循环后slow指针前方的数值都小于自己，后面的都大于自己
func sortList(head *ListNode) *ListNode {
	quickSort(head, nil)
	return head // 返回排序后的链表
}

// 用于交换两个节点上的数值
func swap(a, b *ListNode) {
	a.Val, b.Val = b.Val, a.Val
}

// 快速排序的主体函数
func quickSort(head, end *ListNode) {
	if head == end || head.Next == end {
		return // 头节点已经是结束节点或要排序只有一个头节点直接返回
	}
	pivot := head.Val // 头节点的值作为pivot
	slow, fast := head, head.Next // 定义快慢指针
	for fast != end { // 当快指针不等于结束指针时，不断执行以下操作
		if fast.Val <= pivot { // 如果快指针小于pivot
			fmt.Println(slow.Val)
			fmt.Println(fast.Val)
			printList(head)
			slow = slow.Next // 不断向后移动慢指针
			swap(slow, fast) // 交换快慢指针指向的节点值
			fmt.Println(slow.Val)
			printList(head)
		}
		fast = fast.Next // 快指针向后移动一位
	} // 循环结束后交换头节点和慢指针指向的值,把pivot放在大小两堆数值的中间
	swap(head, slow)
	quickSort(head, slow) // 递归处理pivot左右两边的链表
	quickSort(slow.Next, end)
}

func main()  {
	node:=&ListNode{
		3,&ListNode{4,&ListNode{5,&ListNode{2,&ListNode{1,nil},},},},
	}
	printList(node)
	sortList(node)
	printList(node)

}

func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val)
		fmt.Print("->")
		head = head.Next
	}
	fmt.Println()
}
