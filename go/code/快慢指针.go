package code

// lc 876
func middleNode(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

// lc 141
func hasCycle(head *ListNode) bool {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}

// lc 142
func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			for slow != head {
				slow = slow.Next
				head = head.Next
			}
			return slow
		}
	}
	return nil
}

// lc 143
func reorderList(head *ListNode) {
	mid := middleNode(head)
	head2 := reverseList(mid)
	for head2.Next != nil {
		next := head.Next
		next2 := head2.Next
		head2.Next = head.Next
		head.Next = head2
		head2 = next2
		head = next
	}
}
