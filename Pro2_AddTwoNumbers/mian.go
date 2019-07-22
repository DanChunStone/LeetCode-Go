package main

/*
给出两个非空的链表用来表示两个非负的整数。其中，它们各自的位数是按照逆序的方式存储的，并且它们的每个节点只能存储一位数字。
如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
 */

func main()  {
	p := ListNode{2,nil}
	p.Next = new(ListNode)
	p.Next.Val = 4
	p.Next.Next = new(ListNode)
	p.Next.Next.Val = 3

	q := ListNode{5,nil}
	q.Next = new(ListNode)
	q.Next.Val = 6
	q.Next.Next = new(ListNode)
	q.Next.Next.Val = 4

	addTwoNumbers(&p,&q)
}

//Definition for singly-linked list.
type ListNode struct {
	Val int
	Next *ListNode
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 用于运算的临时变量
	p,q := l1,l2
	var carry int

	// 记录结果的指针变量
	var head,cur *ListNode

	for ; p != nil && q != nil; p,q = p.Next,q.Next {
		if head == nil {
			cur = new(ListNode)
			head = cur
		}else {
			cur.Next = new(ListNode)
			cur = cur.Next
		}

		// 计算当前位与进位并保存
		cur.Val = p.Val + q.Val + carry
		carry = cur.Val / 10
		cur.Val %= 10
	}

	// 复制p/q中多余的部分
	for ; p != nil; p = p.Next {
		cur.Next = new(ListNode)
		cur = cur.Next

		cur.Val = p.Val + carry
		carry = cur.Val / 10
		cur.Val %= 10
	}
	for ; q != nil; q = q.Next {
		cur.Next = new(ListNode)
		cur = cur.Next

		cur.Val = q.Val + carry
		carry = cur.Val / 10
		cur.Val %= 10
	}
	if carry != 0 {
		cur.Next = new(ListNode)
		cur = cur.Next
		cur.Val = carry
	}

	return head
}
