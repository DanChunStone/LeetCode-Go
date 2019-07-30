package main

import "fmt"

func main()  {
	node3 := ListNode{4,nil}
	node2 := ListNode{2,&node3}
	node1 := ListNode{1,&node2}

	node6 := ListNode{4,nil}
	node5 := ListNode{3,&node6}
	node4 := ListNode{1,&node5}

	node8 := ListNode{6,nil}
	node7 := ListNode{2,&node8}

	lists := []*ListNode {&node1,&node4,&node7}
	l := mergeKLists(lists)

	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}
}

// Definition for singly-linked list.
type ListNode struct {
    Val int
    Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	var newLists,head *ListNode

	for !isAllNil(lists) {
		minNode := getMin(lists)
		if newLists == nil {
			newLists = minNode
			head = newLists
		}else {
			newLists.Next = minNode
			newLists = newLists.Next
		}
	}

	return head
}

func isAllNil(lists []*ListNode) bool {
	for _,l := range lists {
		if l != nil {
			return false
		}
	}
	return true
}

func getMin(lists []*ListNode) *ListNode {
	// -1代表min尚未赋值
	min := -1

	// 遍历链表slice
	for i,l := range lists {
		// 若l为空，跳过
		if l != nil {
			if min == -1 { // 若min尚未赋值，直接赋值
				min = i
			}else if l.Val < lists[min].Val { // 若l的值小于min记录的值，更新min
				min = i
			}
		}
	}

	// 获取最小结点，并后移其指针
	minNode := lists[min]
	lists[min] = lists[min].Next
	return minNode
}