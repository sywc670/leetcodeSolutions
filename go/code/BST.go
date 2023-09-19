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

// 450. 删除二叉搜索树中的节点
func deleteNode(root *TreeNode, key int) *TreeNode {
	switch {
	case root == nil:
		return nil
	case root.Val > key:
		root.Left = deleteNode(root.Left, key)
	case root.Val < key:
		root.Right = deleteNode(root.Right, key)
	case root.Left == nil || root.Right == nil:
		if root.Left != nil {
			return root.Left
		}
		return root.Right
	default:
		successor := root.Right
		for successor.Left != nil {
			successor = successor.Left
		}
		successor.Right = deleteNode(root.Right, successor.Val)
		successor.Left = root.Left
		return successor
	}
	return root
}

// 841. 钥匙和房间
func canVisitAllRooms(rooms [][]int) bool {
	visited := make([]bool, len(rooms))
	var dfs func(int)
	dfs = func(i int) {
		visited[i] = true
		for _, key := range rooms[i] {
			if visited[key] {
				continue
			}
			dfs(key)
		}
	}
	dfs(0)

	for _, v := range visited {
		if !v {
			return false
		}
	}
	return true
}

// 399. 除法求值
