package linkedlist

func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	carryOver := 0
	head := &ListNode{
		Val: 0,
	}
	cpy := head
	for l1 != nil || l2 != nil {
		l1Val, l2Val := 0, 0
		if l1 != nil {
			l1Val = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			l2Val = l2.Val
			l2 = l2.Next
		}
		sum := l1Val + l2Val + carryOver
		place := sum % 10
		carryOver = sum / 10
		cpy.Next = &ListNode{
			Val:  place,
			Next: nil,
		}
		cpy = cpy.Next
	}
	if carryOver != 0 {
		cpy.Next = &ListNode{Val: carryOver}
	}

	return head.Next
}
