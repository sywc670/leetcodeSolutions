package code

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

// 1679. K 和数对的最大数目
func maxOperations(nums []int, k int) (count int) {
	m := make(map[int]int)
	for _, n := range nums {
		if m[k-n] > 0 {
			count++
			m[k-n]--
		} else {
			m[n]++
		}
	}
	return
}
func maxOperationsOld(nums []int, k int) (count int) {
	m := make(map[int]int)
	for _, n := range nums {
		_, ok := m[n]
		if !ok {
			m[n] = 1
			continue
		}
		m[n]++
	}
	for key := range m {
		for m[key] > 0 {
			if k == 2*key {
				count += m[key] / 2
				break
			}
			peer := k - key
			if _, ok := m[peer]; ok && m[peer] > 0 {
				count++
				m[peer]--
				m[key]--
			} else {
				break
			}
		}
	}
	return
}

// 1456. 定长子串中元音的最大数目
func maxVowels(s string, k int) (ans int) {
	var cnt int
	for i, in := range s {
		if in == 'a' || in == 'e' || in == 'i' || in == 'o' || in == 'u' {
			cnt++
		}
		if i < k-1 {
			continue
		}
		ans = max(ans, cnt)
		out := s[i-k+1]
		if out == 'a' || out == 'e' || out == 'i' || out == 'o' || out == 'u' {
			cnt--
		}
	}
	return
}

// 1004. 最大连续1的个数 III
// 下面方法是统计1，统计0的更好
func longestOnes(nums []int, k int) (ans int) {
	left, right := 0, 0
	var ones int
	for ; right < len(nums); right++ {
		l := right - left + 1
		if nums[right] == 1 {
			ones++
		}
		if l <= ones+k {
			ans = max(ans, l)
			continue
		}
		for l > ones+k && l > 1 {
			if nums[left] == 1 {
				ones--
			}
			left++
			l = right - left + 1
		}
	}
	return
}

func longestOnesV1(nums []int, k int) (ans int) {
	left, cnt0 := 0, 0
	for right, x := range nums {
		cnt0 += 1 - x
		for cnt0 > k {
			cnt0 -= 1 - nums[left]
			left++
		}
		ans = max(ans, right-left+1)
	}
	return
}

// 27. 移除元素
func removeElement(nums []int, val int) (n int) {
	for _, v := range nums {
		if v != val {
			nums[n] = v
			n++
		}
	}

	return
}

// 977. 有序数组的平方
func sortedSquares(nums []int) []int {
	n := len(nums)
	left, right := 0, n-1
	ans := make([]int, n)

	for i := n - 1; i > -1; i-- {
		x, y := nums[left], nums[right]

		if -x < y {
			ans[i] = y * y
			right--
			continue
		}

		ans[i] = x * x
		left++

	}

	return ans
}

// 394. 字符串解码
func decodeString(s string) string {
	numStack := make([]int, 0)
	strStack := make([]string, 0)
	str := ""
	num := 0

	for _, b := range s {
		if b >= '0' && b <= '9' {
			n, _ := strconv.Atoi(string(b))
			num = 10*num + n

		} else if b == '[' {
			numStack = append(numStack, num)
			strStack = append(strStack, str)
			str = ""
			num = 0

		} else if b == ']' {
			count := numStack[len(numStack)-1]
			numStack = numStack[:len(numStack)-1]
			stacked := strStack[len(strStack)-1]
			strStack = strStack[:len(strStack)-1]

			str = stacked + strings.Repeat(str, count)

		} else {
			str += string(b)
		}
	}
	return str
}

// 649. Dota2 参议院
func predictPartyVictory(senate string) string {
	var radiant, dire []int

	for i, s := range senate {
		if s == 'R' {
			radiant = append(radiant, i)
		} else {
			dire = append(dire, i)
		}
	}

	for len(dire) > 0 && len(radiant) > 0 {
		if radiant[0] < dire[0] {
			radiant = append(radiant, radiant[0]+len(senate))
		} else {
			dire = append(dire, dire[0]+len(senate))
		}
		radiant = radiant[1:]
		dire = dire[1:]
	}

	if len(dire) > 0 {
		return "Dire"
	}
	return "Radiant"
}

// 2095. 删除链表的中间节点
func deleteMiddle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	slow, fast := head, head.Next.Next

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next

	return head
}

// 2130. 链表最大孪生和
func pairSum(head *ListNode) (ans int) {
	slow, fast := head, head.Next
	for fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	// reverse linknode
	// 利用last来保存cur之后要插入的节点，节点为空时刚好last.Next也为空，不用特意赋值
	last := slow.Next
	for last.Next != nil {
		cur := last.Next
		last.Next = cur.Next
		cur.Next = slow.Next
		slow.Next = cur
	}

	// add and compare
	p, q := head, slow.Next
	for q != nil {
		sum := p.Val + q.Val
		ans = max(ans, sum)
		p = p.Next
		q = q.Next
	}

	return
}

// 1448. 统计二叉树中好节点的数目
func goodNodes(root *TreeNode) (cnt int) {
	var dfs func(*TreeNode, int)
	dfs = func(tn *TreeNode, pathMax int) {
		if tn == nil {
			return
		}

		if tn.Val >= pathMax {
			cnt++
			pathMax = tn.Val
		}

		dfs(tn.Left, pathMax)
		dfs(tn.Right, pathMax)
	}

	dfs(root, math.MinInt)
	return
}

// 1372. 二叉树中的最长交错路径
func longestZigZag(root *TreeNode) (ans int) {
	var dfs func(*TreeNode, bool, int)
	dfs = func(node *TreeNode, left bool, n int) {
		if node == nil {
			return
		}

		ans = max(ans, n)

		switch left {
		case true:
			dfs(node.Left, true, 1)
			dfs(node.Right, false, n+1)
		case false:
			dfs(node.Right, false, 1)
			dfs(node.Left, true, n+1)
		}
	}
	dfs(root.Left, true, 1)
	dfs(root.Right, false, 1)

	return
}

// 114. 二叉树展开为链表
// 这里用先序遍历，先将左右子树变成链表，再接上
func flatten(node *TreeNode) {
	if node == nil {
		return
	}
	flatten(node.Left)
	flatten(node.Right)

	tmp := node.Right
	node.Right = node.Left
	node.Left = nil
	p := node
	for p.Right != nil {
		p = p.Right
	}
	p.Right = tmp
}

func flattenV1(node *TreeNode) {
	flattenWithTail(node)
}

func flattenWithTail(node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}

	leftTail := flattenWithTail(node.Left)
	rightTail := flattenWithTail(node.Right)

	// 如果左子树存在，把它插入到右边，然后接上原来的右子树
	if node.Left != nil {
		right := node.Right
		node.Right = node.Left
		node.Left = nil
		leftTail.Right = right
	}

	if rightTail != nil {
		return rightTail
	}

	if leftTail != nil {
		return leftTail
	}

	return node
}

// 48. 旋转图像
func rotate(matrix [][]int) {
	n := len(matrix)
	// 水平翻转
	for i := 0; i < n/2; i++ {
		matrix[i], matrix[n-1-i] = matrix[n-1-i], matrix[i]
	}
	// 主对角线翻转
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}

// 2. 两数相加
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	p1, p2, p := l1, l2, head
	var carry int
	for p1 != nil || p2 != nil {
		n := &ListNode{}
		v := carry
		if p1 != nil {
			v += p1.Val
			p1 = p1.Next
		}
		if p2 != nil {
			v += p2.Val
			p2 = p2.Next
		}
		carry = 0
		if v >= 10 {
			carry = 1
			v = v - 10
		}
		n.Val = v
		p.Next = n
		p = p.Next
	}
	if carry > 0 {
		p.Next = &ListNode{Val: carry}
	}
	return head.Next
}

// 21. 合并两个有序链表
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{}
	p1, p2, p := list1, list2, head
	for p1 != nil && p2 != nil {
		if p1.Val < p2.Val {
			p.Next = p1
			p1 = p1.Next
		} else {
			p.Next = p2
			p2 = p2.Next
		}
		p = p.Next
	}
	if p1 != nil {
		p.Next = p1
	}
	if p2 != nil {
		p.Next = p2
	}
	return head.Next
}

// 160. 相交链表
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	length := func(head *ListNode) (l int) {
		p := head
		for p != nil {
			p = p.Next
			l++
		}
		return
	}
	la, lb := length(headA), length(headB)
	pa, pb := headA, headB
	// 找到更长的先走
	first, last, diff := pa, pb, la-lb
	if lb > la {
		first, last, diff = pb, pa, lb-la
	}
	for diff > 0 {
		first = first.Next
		diff--
	}
	// 一起走
	for first != nil {
		if first == last {
			return first
		}
		first, last = first.Next, last.Next
	}
	return nil
}

// https://leetcode.cn/problems/intersection-of-two-linked-lists/solutions/2958778/tu-jie-yi-zhang-tu-miao-dong-xiang-jiao-m6tg1
// 算法题图片：160.png
// 第二种思路：(a+c)+b = (b+c)+a
func getIntersectionNodeV1(headA, headB *ListNode) *ListNode {
	p, q := headA, headB
	for p != q {
		if p != nil {
			p = p.Next
		} else {
			p = headB
		}
		if q != nil {
			q = q.Next
		} else {
			q = headA
		}
	}
	return p
}

// 240. 搜索二维矩阵 II
// 二分
func searchMatrix(matrix [][]int, target int) bool {
	for _, row := range matrix {
		i := sort.SearchInts(row, target)
		if i < len(row) && row[i] == target {
			return true
		}
	}
	return false
}

// 排除法
// 与右上角比较,从而进行排除
func searchMatrixV1(matrix [][]int, target int) bool {
	i, j := 0, len(matrix[0])-1
	for i < len(matrix) && j >= 0 {
		if matrix[i][j] == target {
			return true
		}
		if matrix[i][j] > target {
			j--
		} else {
			i++
		}
	}
	return false
}

// 24. 两两交换链表中的节点
// 递归做法
func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	var reverse func(*ListNode)
	reverse = func(head *ListNode) {
		p := head.Next
		if p != nil && p.Next != nil {
			q := p.Next
			p.Next = q.Next
			q.Next = p
			head.Next = q
			reverse(p)
		}
	}
	reverse(dummy)
	return dummy.Next
}
