package main

import (
	"fmt"
	"leetcode100/problem"
)

func main() {
	//var a, b [26]int
	//a[0] = 2
	//a[1] = 2
	//
	//b[0] = 1
	//b[1] = 2
	//
	//fmt.Println(a == b)

	//problem.FindAnagrams("cbaebabacd", "abc")
	//problem.MaxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})

	n1 := &problem.ListNode{Val: 1}
	n2 := &problem.ListNode{Val: 2}
	n3 := &problem.ListNode{Val: 3}
	n4 := &problem.ListNode{Val: 4}
	n5 := &problem.ListNode{Val: 5}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5
	fmt.Println(problem.ReverseKGroup(n1, 2))

}
