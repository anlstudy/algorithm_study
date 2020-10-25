package main

import "fmt"

// 删除链表倒数第N个节点


//给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。
//
//示例：
//
//给定一个链表: 1->2->3->4->5, 和 n = 2.
//
//当删除了倒数第二个节点后，链表变为 1->2->3->5.
//说明：
//
//给定的 n 保证是有效的。
//
//进阶：
//
//你能尝试使用一趟扫描实现吗？
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。


type ListNode struct {
     Val int
     Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
     if head == nil {
          return nil
     }

     length:=1
     first,second := head,head
     for first.Next!=nil{
          length++
          first = first.Next
     }
     if length == n{
          return head.Next
     }
     del:=length-n-1

     for del >0{
          del --
          second = second.Next
     }

     second.Next = second.Next.Next

     return  head
}

func removeNthFromEnd2(head *ListNode, n int) *ListNode {
     result := &ListNode{}
     result.Next = head
     var pre *ListNode
     cur := result
     i := 1
     for head != nil {
          if i >= n {
               pre = cur
               cur = cur.Next
          }
          head = head.Next
          i++
     }
     pre.Next = pre.Next.Next
     return result.Next
}

func main(){
     head:=&ListNode{
          Val:1,
          Next:&ListNode{
               Val:2,
               Next:&ListNode{
                    Val:  3,
                    Next: nil,
               },
     },
     }

     printList(removeNthFromEnd(head,3))
}


func printList(head *ListNode){
     for head!=nil{
          fmt.Println(head.Val)
          head = head.Next
     }
}