package code

type ListNode struct {
	Val  int
	Next *ListNode
}

// trick: dummynode or sentry

// lc 206
// 头插法
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

// 递归
// 就是把链表的每个指针反转，在反转当前的节点指针前，把后续链表节点给反转了，把尾节点信息返回作为反转后的头结点
func reverseListV2(head *ListNode) *ListNode {
	var reverse func(*ListNode, *ListNode) *ListNode
	reverse = func(cur, pre *ListNode) *ListNode {
		if cur == nil {
			// 终止条件，返回反转后的头结点，其实就是反转前最后一个节点
			return pre
		}
		res := reverse(cur.Next, cur)
		// 当前操作，将当前节点指向上一个节点以反转
		cur.Next = pre
		return res
	}
	return reverse(head, nil)
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

// 328. 奇偶链表
// 未掌握
// solve:巧妙的点在于每次走了两步，但是判断是否为空却不是那两步，最后一步放在下一次循环里面进行判断
func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	evenhead := head.Next
	eventail := evenhead
	oddtail := head
	for eventail != nil && eventail.Next != nil {
		oddtail.Next = eventail.Next
		oddtail = oddtail.Next
		eventail.Next = oddtail.Next
		eventail = eventail.Next
	}
	oddtail.Next = evenhead
	return head
}
