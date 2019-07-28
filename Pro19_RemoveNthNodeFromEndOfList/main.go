package main

import "fmt"

func main()  {
	node5 := ListNode{5,nil}
	node4 := ListNode{4,&node5}
	node3 := ListNode{3,&node4}
	node2 := ListNode{2,&node3}
	node1 := ListNode{1,&node2}
	node1.print()

	node := removeNthFromEnd(&node1,5)
	node.print()
}

// Definition for singly-linked list.
type ListNode struct {
    Val int
    Next *ListNode
}
// 输出链表
func(l *ListNode) print() {
	if l == nil {
		fmt.Println("nil")
		return
	}

	head := l
	for ; head.Next != nil; head = head.Next {
		fmt.Print(head.Val)
		fmt.Print(" -> ")
	}
	fmt.Println(head.Val)
}

// 删除链表的倒数第n个结点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || n == 0{
		return nil
	}

	// 采用先后指针，first先走n个结点，next再开始走，直到first到达尾部
	first,next := head,head

	// first向前走n个结点
	for n > 0 && first != nil {
		first = first.Next
		n--
	}
	// 如果链表长度不足n，则返回
	if n > 0 {
		return head
	}
	// 如果链表长度恰好为n，则删除头结点
	if first == nil {
		return head.Next
	}

	// 让first多走一步，使得next指向该删除的结点的前一个，便于删除
	first = first.Next
	for first != nil {
		first = first.Next
		next = next.Next
	}
	// 删除next此时的后一个结点
	next.Next = next.Next.Next

	return head
}