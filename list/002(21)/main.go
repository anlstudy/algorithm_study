package main

import "fmt"

/*
第21题：合并两个有序链表
将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
示例：

输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	prehead := &ListNode{}
	result := prehead
	for l1!=nil && l2!=nil{
		if l1.Val <=l2.Val {
			prehead.Next = l1
			l1 = l1.Next
		}else{
			prehead.Next = l2
			l2 = l2.Next
		}
		prehead = prehead.Next
	}

	if l1!=nil {
		prehead.Next = l1
	}
	if l2!=nil {
		prehead.Next = l2
	}
	return result.Next
}

func main() {

	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	l2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	printList(l1)
	printList(l2)
	printList(mergeTwoLists(l1,l2))
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val)
		fmt.Print("->")
		head = head.Next
	}
	fmt.Println()
}
