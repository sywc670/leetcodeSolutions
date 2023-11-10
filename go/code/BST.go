package code

import "math"

// lc 98
func isValidBST(root *TreeNode) bool {
	var dfs func(node *TreeNode, left, right int) bool
	dfs = func(node *TreeNode, left, right int) bool {
		if node == nil {
			return true
		}
		if node.Val <= left || node.Val >= right {
			return false
		}
		return dfs(node.Left, left, node.Val) && dfs(node.Right, node.Val, right)
	}
	return dfs(root, math.MinInt, math.MaxInt)
}
func isValidBSTV2(root *TreeNode) bool {
	preVal := math.MinInt
	var dfs func(node *TreeNode) bool
	dfs = func(node *TreeNode) bool {
		if node == nil {
			return true
		}
		l := dfs(node.Left)
		if node.Val <= preVal {
			return false
		}
		preVal = node.Val
		r := dfs(node.Right)
		return l && r
	}
	return dfs(root)
}

// 700. 二叉搜索树中的搜索
func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	l := searchBST(root.Left, val)
	if l != nil {
		return l
	}
	return searchBST(root.Right, val)
}
