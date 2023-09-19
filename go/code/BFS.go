package code

// lc 102
func levelOrder(root *TreeNode) (ans [][]int) {
	if root == nil {
		return
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		tmp := q
		q = nil
		res := make([]int, 0)
		for _, node := range tmp {
			res = append(res, node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		ans = append(ans, res)
	}
	return
}

func levelOrderV2(root *TreeNode) (ans [][]int) {
	if root == nil {
		return
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		n := len(q)
		res := make([]int, n)
		for i := 0; i < n; i++ {
			node := q[0]
			q = q[1:]
			res[i] = node.Val
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		ans = append(ans, res)
	}
	return
}

// lc 103
func zigzagLevelOrder(root *TreeNode) (ans [][]int) {
	if root == nil {
		return
	}
	q := []*TreeNode{root}
	for even := false; len(q) > 0; even = !even {
		n := len(q)
		vals := make([]int, n) // 大小已知
		for i := 0; i < n; i++ {
			node := q[0]
			q = q[1:]
			if even {
				vals[n-1-i] = node.Val // 倒着添加
			} else {
				vals[i] = node.Val
			}
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		ans = append(ans, vals)
	}
	return
}

// lc 513
func findBottomLeftValue(root *TreeNode) int {
	node := root
	q := []*TreeNode{root}
	for len(q) > 0 {
		node, q = q[0], q[1:]
		if node.Right != nil {
			q = append(q, node.Right)
		}
		if node.Left != nil {
			q = append(q, node.Left)
		}
	}
	return node.Val
}
