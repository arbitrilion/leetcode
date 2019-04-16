package addtwonumbers

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	p1, p2 := l1, l2
	carrybit := 0
	var resultList ListNode
	pr := &resultList
	carrybit, result := addSum(p1.Val, p2.Val, carrybit)
	pr.Val = result

	for {
		var resultItem ListNode

		if p1.Next != nil && p2.Next != nil {
			pr.Next = &resultItem
			p1, p2, pr = p1.Next, p2.Next, pr.Next
			carrybit, pr.Val = addSum(p1.Val, p2.Val, carrybit)
		} else if p1.Next == nil && p2.Next != nil {
			pr.Next = &resultItem
			p2, pr = p2.Next, pr.Next
			carrybit, pr.Val = addSum(0, p2.Val, carrybit)
		} else if p2.Next == nil && p1.Next != nil {
			pr.Next = &resultItem
			p1, pr = p1.Next, pr.Next
			carrybit, pr.Val = addSum(0, p1.Val, carrybit)
		} else if p2.Next == nil && p1.Next == nil {
			if carrybit == 0 {
				pr.Next = nil
			} else {
				pr.Next = &resultItem
				pr = pr.Next
				pr.Val = carrybit
			}
			break
		}
	}
	return &resultList
}

func addSum(v1 int, v2 int, v3 int) (carrybit int, result int) {
	addsum := v1 + v2 + v3
	if addsum < 10 {
		carrybit = 0
		result = addsum
	} else {
		carrybit = 1
		result = addsum - 10
	}
	return carrybit, result
}
