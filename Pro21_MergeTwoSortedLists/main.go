package main

import "fmt"

func main()  {
	node3 := ListNode{4,nil}
	node2 := ListNode{2,&node3}
	node1 := ListNode{1,&node2}

	node6 := ListNode{4,nil}
	node5 := ListNode{3,&node6}
	node4 := ListNode{1,&node5}

	l := mergeTwoLists(&node1,&node4)

	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}
}

//Definition for singly-linked list.
type ListNode struct {
    Val int
    Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}

	list := make([]ListNode,0)
	h1,h2 := l1,l2

	// 遍历两链表，合并都为空的部分
	for h1 != nil && h2 != nil {
		if h1.Val <= h2.Val {
			list = append(list,*h1)
			h1 = h1.Next
		}else {
			list = append(list, *h2)
			h2 = h2.Next
		}
	}
	// 清除剩余部分
	for h1 != nil {
		list = append(list,*h1)
		h1 = h1.Next
	}
	for h2 != nil {
		list = append(list,*h2)
		h2 = h2.Next
	}

	// 重新构造链表结构
	for i := 1; i < len(list); i++ {
		list[i-1].Next = &list[i]
	}
	return &list[0]
}
