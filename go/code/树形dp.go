package code

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// lc 543
// 二叉树求最大值类型的题可能会用到遍历每个节点比较存储最大值的方法
// 在顶点求链的值用加法，而子树求链的值用max比较，因为子树的链要过顶点必须不能分叉
// 原问题：root为根的二叉树求最大直径
// 子问题：node为根的二叉树求最大直径，最大直径经过node
// 下一个子问题：node的左右子树为根的二叉树求最大直径，最大直径经过node左右子树
// 当前操作：拿到node左右子树的最大深度，由于是后序遍历，所以深度放在返回值里
// 返回node的最大深度，并记录经过node的最大直径
func diameterOfBinaryTree(root *TreeNode) int {
	var ans int
	var dfs func(*TreeNode) int
	dfs = func(t *TreeNode) int {
		if t == nil {
			return -1
		}
		l := dfs(t.Left)
		r := dfs(t.Right)
		// return value is depth
		ans = max(ans, l+r+2)
		return max(l, r) + 1
	}
	dfs(root)
	return ans
}

// lc 124
// If the result of subtree is below zero
// return zero because that result is not useful
// 子问题：以node为根的二叉树的最大路径和，经过node
// 下一个子问题：以node左右子树为根的二叉树的最大路径和，经过node左右子树
// 当前操作：拿到node左右子树的返回值，计算最大路径和并记录，返回单边最大路径和
func maxPathSum(root *TreeNode) int {
	ans := math.MinInt
	var dfs func(*TreeNode) int
	dfs = func(t *TreeNode) (res int) {
		if t == nil {
			return 0
		}
		l := dfs(t.Left)
		r := dfs(t.Right)
		res = max(l, r) + t.Val
		ans = max(ans, l+r+t.Val)
		if res < 0 {
			return 0
		}
		return
	}
	dfs(root)
	return ans
}

// lc 2246
// 子问题：以node为根的树的最长不相同路径，经过node
// 下一个子问题：以node子树为根的树的最长不相同路径，经过node左右子树
// 当前操作：选出最长和次长的子树，计算单边最长不相同路径，返回给上一个节点，比较经过该node的最长不相同路径
func longestPath(parent []int, s string) (ans int) {
	n := len(parent)
	g := make([][]int, n)
	for i := 1; i < n; i++ {
		pa := parent[i]
		g[pa] = append(g[pa], i)
	}

	var dfs func(int) int
	dfs = func(x int) (maxLen int) {
		for _, y := range g[x] {
			len := dfs(y) + 1
			// 需要先迭代再进行判断，否则可能会漏掉
			if s[y] != s[x] {
				ans = max(ans, maxLen+len)
				maxLen = max(maxLen, len)
			}
		}
		return
	}
	dfs(0)
	return ans + 1
}

// lc 337
// 打家劫舍
// 思路：由于不能连续偷取，所以需要考虑一个节点选与不选，可以添加一个状态表示选与不选
// 但是由于树形dp是需要后序遍历来做的，所以这个状态不能加在参数里，而是在返回两个值
func robV(root *TreeNode) (ans int) {
	var dfs func(*TreeNode) (int, int)
	dfs = func(node *TreeNode) (int, int) {
		if node == nil {
			return 0, 0
		}
		lRob, lNotRob := dfs(node.Left)
		rRob, rNotRob := dfs(node.Right)
		notRob := max(lRob, lNotRob) + max(rRob, rNotRob)
		return lNotRob + rNotRob + node.Val, notRob
	}
	return max(dfs(root))
}

// lc 968
// 根据摄像头的位置来分类讨论
// note: 未完全掌握
func minCameraCover(root *TreeNode) int {
	var dfs func(node *TreeNode) (int, int, int)
	dfs = func(node *TreeNode) (int, int, int) {
		if node == nil {
			return math.MaxInt / 2, 0, 0 // 除 2 防止加法溢出
		}
		lChoose, lByFa, lByChildren := dfs(node.Left)
		rChoose, rByFa, rByChildren := dfs(node.Right)
		choose := min(lChoose, lByFa) + min(rChoose, rByFa) + 1
		byFa := min(lChoose, lByChildren) + min(rChoose, rByChildren)
		byChildren := min(min(lChoose+rByChildren, lByChildren+rChoose), lChoose+rChoose)
		return choose, byFa, byChildren
	}
	choose, _, byChildren := dfs(root)
	return min(choose, byChildren)
}
