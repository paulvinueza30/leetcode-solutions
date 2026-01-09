package linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(h *ListNode) bool {
	if h == nil || h.Next == nil {
		return false
	}
	slow, fast := h, h.Next

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}

	return false
}
