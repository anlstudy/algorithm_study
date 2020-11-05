package main

import "fmt"

//给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
//
//如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
//
//您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
//
//示例：
//
//输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
//输出：7 -> 0 -> 8
//原因：342 + 465 = 807

type ListNode struct {
	Val  int
	Next *ListNode
}


//func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
//	list := &ListNode{0, nil}
//	//这里用一个result，只是为了后面返回节点方便，并无他用
//	result := list
//	tmp := 0
//	for l1 != nil || l2 != nil || tmp != 0 {
//		if l1 != nil {
//			tmp += l1.Val
//			l1 = l1.Next
//		}
//		if l2 != nil {
//			tmp += l2.Val
//			l2 = l2.Next
//		}
//		list.Next = &ListNode{tmp % 10, nil}
//		tmp = tmp / 10
//		list = list.Next
//	}
//	return result.Next
//}


func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 ==nil && l2==nil {
		return nil
	}
	res:=&ListNode{
		Val: 0,
		Next: nil,
	}
	head:=res
	ext:=0
	for l1!=nil && l2!=nil{
		head.Next = &ListNode{(l1.Val+l2.Val+ext)%10, nil}
		ext = (l1.Val+l2.Val+ext)/10
		head = head.Next
		l1 = l1.Next
		l2 = l2.Next
	}

	for l1!=nil{
		head.Next = &ListNode{(l1.Val+ext)%10, nil}
		ext = (l1.Val+ext)/10
		head = head.Next
		l1 = l1.Next
	}

	for l2!=nil{
		head.Next = &ListNode{(l2.Val+ext)%10, nil}
		ext = (l2.Val+ext)/10
		head = head.Next
		l2 = l2.Next
	}

	if ext!=0{
		head.Next = &ListNode{ext, nil}
	}
	return res.Next
}

func main(){
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
	printList(addTwoNumbers(l1,l2))
}



func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val)
		fmt.Print("->")
		head = head.Next
	}
	fmt.Println()
}
