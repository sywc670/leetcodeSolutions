package code

type ListNode struct {
	Val  int
	Next *ListNode
}

// trick: dummynode or sentry

// lc 206
func reverseList(head *ListNode) *ListNode {
	h := new(ListNode)
	for p := head; p != nil; {
		tmp := p.Next
		p.Next = h.Next
		h.Next = p
		p = tmp
	}
	return h.Next
}

// lc 92
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == right {
		return head
	}
	dummy := &ListNode{Next: head}
	insert := dummy
	for count := left; count > 1; count-- {
		insert = insert.Next
	}
	cur := insert.Next
	// 将尾部区域接在头部区域的后面，中间反转区域相当于单独链表
	tail := cur
	for count := right - left + 1; count > 0; count-- {
		tail = tail.Next
	}
	insert.Next = tail
	for count := right - left + 1; count > 0; count-- {
		tmp := cur.Next
		cur.Next = insert.Next
		insert.Next = cur
		cur = tmp
	}

	return dummy.Next
}
func reverseBetweenV2(head *ListNode, left int, right int) *ListNode {
	if left == right {
		return head
	}
	dummy := &ListNode{Next: head}
	insert := dummy
	for count := left; count > 1; count-- {
		insert = insert.Next
	}
	cur := insert.Next
	last := cur // last是反转区域第一个节点，每次更新其值为cur，最后会指向尾部区域的第一个节点
	for count := right - left + 1; count > 0; count-- {
		tmp := cur.Next
		cur.Next = insert.Next
		insert.Next = cur
		cur = tmp
		last.Next = cur
	}
	return dummy.Next
}

// lc 25
func reverseKGroup(head *ListNode, k int) *ListNode {
	// TODO: 用上一题的代码修改
	return head
}
