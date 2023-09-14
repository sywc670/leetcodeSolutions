package code

// lc 104
func maxDepth(root *TreeNode) int {
	var dfs func(*TreeNode) int
	dfs = func(n *TreeNode) int {
		if n == nil {
			return 0
		}
		l := dfs(n.Left)
		r := dfs(n.Right)
		return max(l, r) + 1
	}
	return dfs(root)
}

// lc 100
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p != nil && q != nil {
		if p.Val == q.Val {
			return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
		}
	}
	return false
}
func isSameTreeV2(p, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}
	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

// lc 101
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var dfs func(p, q *TreeNode) bool
	dfs = func(p, q *TreeNode) bool {
		if p == nil || q == nil {
			if p == nil && q == nil {
				return true
			}
			return false
		}
		return p.Val == q.Val && dfs(p.Left, q.Right) && dfs(p.Right, q.Left)
	}
	return dfs(root.Left, root.Right)
}
func isSymmetricV2(root *TreeNode) bool {
	return isSameTreeV2(root.Left, root.Right)
}

// lc 110
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	sub := maxDepth(root.Left) - maxDepth(root.Right)
	if sub == -1 || sub == 1 || sub == 0 {
		return isBalanced(root.Left) && isBalanced(root.Right)
	}
	return false
}
