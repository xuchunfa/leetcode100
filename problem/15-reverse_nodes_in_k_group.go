package problem

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	cur := head
	for i := 0; i < k; i++ {
		if cur == nil {
			return head
		}
		cur = cur.Next
	}
	first := cur
	//fmt.Println(first.Val)
	cur = head
	var prev *ListNode
	for cur != nil && cur != first {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	res := prev
	for prev.Next != nil {
		prev = prev.Next
	}
	prev.Next = ReverseKGroup(first, k)
	fmt.Println(res.Val)
	return res
}
