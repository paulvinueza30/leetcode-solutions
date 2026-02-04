package linkedlist

func deleteDuplicates(head *ListNode) *ListNode {
	curr := head

	for curr != nil {
		iter := curr
		for iter != nil && iter.Val == curr.Val {
			iter = iter.Next
		}
		curr.Next = iter
		curr = curr.Next
	}

	return head
}
