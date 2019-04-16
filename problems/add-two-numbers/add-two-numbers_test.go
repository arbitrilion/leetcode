package addtwonumbers

import (
	"fmt"
	"testing"
)

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

var (
	v1, v2 = 789, 465
)

func TestAddTwoNumbers(t *testing.T) {
	l1 := turnIntToReverseArray(v1)
	for i := 0; i < 10; i++ {
		if l1 == nil {
			break
		}
		fmt.Printf("%d ", l1.Val)
		l1 = l1.Next
	}
	fmt.Println("l1.")

	l2 := turnIntToReverseArray(v2)
	for i := 0; i < 10; i++ {
		if l2 == nil {
			break
		}
		fmt.Printf("%d ", l2.Val)
		l2 = l2.Next
	}
	fmt.Println("l2.")

	result := addTwoNumbers(turnIntToReverseArray(v1), turnIntToReverseArray(v2))
	// fmt.Println(result)
	for i := 0; i < 10; i++ {
		if result == nil {
			break
		}
		fmt.Printf("%d ", result.Val)
		result = result.Next
	}
	fmt.Println("result.")

	fmt.Println()
	sum := v1 + v2
	reverseSum := turnIntToReverseArray(sum)
	fmt.Println("Sum: ", sum, reverseSum)
}

func turnIntToReverseArray(val int) *ListNode {
	var arr []int
	for {
		if val < 10 {
			arr = append(arr, val)
			break
		} else {
			arr = append(arr, val%10)
			val = val / 10
		}
	}
	var res []int
	for i := 0; i < len(arr); i++ {
		res = append(res, arr[len(arr)-1-i])
	}
	return makeLinkList(res)
}

var l0 ListNode
var tmp *ListNode

// 注意值传递和指针传递的区别
// 直接传值会导致死循环
func makeLinkList(s []int) *ListNode {
	// fmt.Println("s:", s)
	for k, v := range s {
		if k == 0 {
			l0 = ListNode{v, nil}
			// fmt.Println("l0: ", l0)
			tmp = &l0
		} else {
			var lnew ListNode
			lnew.Val = v
			lnew.Next = tmp
			// fmt.Println("lnew: ", lnew)
			tmp = &lnew
		}
	}
	result := tmp
	// for i := 0; i < 10; i++ {
	// 	fmt.Printf("%d ", tmp.Val)
	// 	tmp = tmp.Next
	// 	if tmp == nil {
	// 		break
	// 	}
	// }
	// fmt.Println("reverse.")
	// fmt.Println()
	return result
}
