package main

import "fmt"

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// 9, 8, 7
	// 6, 5, 4
	l3 := ListNode{0, nil}
	// l2 := ListNode{8, &l3}
	// l1 := ListNode{9, &l2}
	l6 := ListNode{4, nil}
	l5 := ListNode{5, &l6}
	l4 := ListNode{6, &l5}
	resultList := addTwoNumbers(&l4, &l3)
	fmt.Println("resultList:")
	for {
		fmt.Println(resultList.Val)
		resultList = resultList.Next
		if resultList == nil {
			break
		}
	}
}
