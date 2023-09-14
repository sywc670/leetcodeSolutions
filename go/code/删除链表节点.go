package code

// lc 19
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head, Val: -1}
	predel, tail := dummy, dummy
	for count := n; count > 0; count-- {
		tail = tail.Next
	}
	for tail.Next != nil {
		predel = predel.Next
		tail = tail.Next
	}
	predel.Next = predel.Next.Next
	return dummy.Next
}

// lc 83
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	dummy := &ListNode{Next: head, Val: head.Val}
	pre, p := dummy, head
	for p != nil {
		if pre.Val == p.Val {
			pre.Next = p.Next
			p = pre.Next
		} else {
			p = p.Next
			pre = pre.Next
		}
	}
	return dummy
}

// lc 82
func deleteDuplicatesV1(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	cur := dummy
	for cur.Next != nil && cur.Next.Next != nil {
		val := cur.Next.Val
		if cur.Next.Next.Val == val {
			for cur.Next != nil && cur.Next.Val == val {
				cur.Next = cur.Next.Next
			}
		} else {
			cur = cur.Next
		}
	}
	return dummy.Next
}
