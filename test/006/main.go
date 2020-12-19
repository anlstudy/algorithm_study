package main

import (
	"fmt"
	"strconv"
)

//
//题目描述：一个链表，奇数位升序偶数位降序，让链表变成升序的。
//
//比如：1 8 3 6 5 4 7 2 9，最后输出1 2 3 4 5 6 7 8 9

type ListNode struct {
	Val int
	Next *ListNode
}

func main(){
	src:=&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 8,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 6,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val: 4,
							Next: nil,
						},
					},
				},
			},
		},
	}
	printList(src)
	println()
	split_res:= split(src)
	printList(split_res[0])
	println()
	printList(split_res[1])
	println()
	printList(merge(split_res[0],reverse(split_res[1])))
}


func split(list *ListNode) []*ListNode{
	l1:=&ListNode{
		Val: -1,
		Next: nil,
	}
	l2:=&ListNode{
		Val: -1,
		Next: nil,
	}
	head1,head2:=l1,l2
	index:=1
	for list!=nil{
		if index%2 == 1 {
			head1.Next = &ListNode{
				list.Val,
				nil,
			}
			head1 = head1.Next
		}else{
			head2.Next = &ListNode{
				list.Val,
				nil,
			}
			head2 = head2.Next
		}
		list = list.Next
		index++
	}
	return []*ListNode{
		l1.Next,l2.Next,
	}
}


func reverse(list *ListNode) *ListNode{
	if list == nil {
		return list
	}
	var pre,cur *ListNode

	for list!=nil{
		cur = list
		list = list.Next
		cur.Next = pre
		pre = cur
	}

	return pre
}

func merge(l1,l2 *ListNode)*ListNode{
	list:=&ListNode{}
	res:=list
	for l1!=nil && l2!=nil {
		if l1.Val < l2.Val {
			list.Next = l1
			l1 = l1.Next
		}else{
			list.Next = l2
			l2 = l2.Next
		}
		list = list.Next
	}

	for l1!=nil {
		list.Next = l1
		l1 = l1.Next
		list = list.Next
	}

	for l2!=nil {
		list.Next = l2
		l2 = l2.Next
		list = list.Next
	}
	return  res.Next

}


func printList(head *ListNode){
	for head!=nil{
		fmt.Print(strconv.Itoa(head.Val)+"->")
		head = head.Next
	}
}

