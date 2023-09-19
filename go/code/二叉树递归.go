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
func isBalancedV2(root *TreeNode) bool {
	var getHeight func(node *TreeNode) int
	getHeight = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		l := getHeight(node.Left)
		r := getHeight(node.Right)
		if l == -1 || r == -1 {
			return -1
		}
		if sub := l - r; sub <= 1 && sub >= -1 {
			return max(l, r) + 1
		}
		return -1
	}
	if getHeight(root) != -1 {
		return true
	}
	return false
}

// 199. 二叉树的右视图
// DFS
func rightSideViewV1(root *TreeNode) (ans []int) {
	var maxDepth int
	var dfs func(node *TreeNode, maxDepth int)
	dfs = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		if depth > maxDepth {
			ans = append(ans, node.Val)
			maxDepth++
		}
		dfs(node.Right, depth+1)
		dfs(node.Left, depth+1)
	}
	dfs(root, 1)
	return
}

// lc 235
func lowestCommonAncestorV1(root, p, q *TreeNode) *TreeNode {
	v := root.Val
	if p.Val > v && q.Val > v {
		return lowestCommonAncestorV1(root.Right, p, q)
	}
	if p.Val < v && q.Val < v {
		return lowestCommonAncestorV1(root.Left, p, q)
	}
	return root
}

// 236. 二叉树的最近公共祖先
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	return root
}
