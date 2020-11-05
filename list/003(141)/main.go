package main

import "fmt"

/*
给定一个链表，判断链表中是否有环。

如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。
如果 pos 是 -1，则在该链表中没有环。注意：pos 不作为参数进行传递，仅仅是为了标识链表的实际情况。

如果链表中存在环，则返回 true 。 否则，返回 false 。


*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val)
		fmt.Print("->")
		head = head.Next
	}
	fmt.Println()
}

// hash法
func hasCycle2(head *ListNode) bool {
	m:=make(map[*ListNode] int)
	for head!=nil{
		if _,exist:=m[head];exist{
			return true
		}
		m[head] = 1
		head = head.Next
	}
	return false
}



func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	fast := head.Next       // 快指针，每次走两步
	for fast != nil && head != nil && fast.Next != nil {
		if fast == head {   // 快慢指针相遇，表示有环
			return true
		}
		fast = fast.Next.Next
		head = head.Next        // 慢指针，每次走一步
	}
	return false
}

func main(){

}
