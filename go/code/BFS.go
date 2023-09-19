package code

import "math"

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

// 199. 二叉树的右视图
// BFS
func rightSideView(root *TreeNode) (ans []int) {
	if root == nil {
		return
	}
	queue := []*TreeNode{root}
	// 临时栈用于存放遍历节点的左右子节点，直到队列为空
	tmpQ := []*TreeNode{}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node.Left != nil {
			tmpQ = append(tmpQ, node.Left)
		}
		if node.Right != nil {
			tmpQ = append(tmpQ, node.Right)
		}

		if len(queue) == 0 {
			ans = append(ans, node.Val)
			queue = tmpQ
			tmpQ = []*TreeNode{}
		}
	}
	return
}

// 1161. 最大层内元素和
// BFS，以每一个节点来循环，也可以在外层循环以一层来循环，内层以一个来循环，参考下面
func maxLevelSum(root *TreeNode) (ans int) {
	if root == nil {
		return
	}
	queue := []*TreeNode{root}
	// 临时栈用于存放遍历节点的左右子节点，直到队列为空
	tmpQ := []*TreeNode{}
	level, levelSum, maxSum := 1, 0, math.MinInt

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		levelSum += node.Val

		if node.Left != nil {
			tmpQ = append(tmpQ, node.Left)
		}
		if node.Right != nil {
			tmpQ = append(tmpQ, node.Right)
		}

		if len(queue) == 0 {
			if levelSum > maxSum {
				maxSum = levelSum
				ans = level
			}
			levelSum = 0
			level++
			queue = tmpQ
			tmpQ = []*TreeNode{}
		}
	}
	return
}

// 这种方法相当于多了个中间层来处理一层遍历结束的情况，不用在放在遍历节点的逻辑中来处理，可以解耦逻辑
// 并且用的bfs方法也更好，pop是从tmp队列，而push是到q队列，这样判断条件更简单清晰
func maxLevelSumV1(root *TreeNode) int {
	ans, maxSum := 1, root.Val
	q := []*TreeNode{root}
	for level := 1; len(q) > 0; level++ {
		tmp := q
		q = nil
		sum := 0
		for _, node := range tmp {
			sum += node.Val
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		if sum > maxSum {
			ans, maxSum = level, sum
		}
	}
	return ans
}

// 1926. 迷宫中离入口最近的出口
// 技巧：用增量来代替取值；提前判断是否越界
func nearestExit(maze [][]byte, entrance []int) (step int) {
	type pos struct {
		x, y int
	}
	delta := []pos{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	m, n := len(maze), len(maze[0])
	entry := pos{entrance[0], entrance[1]}
	maze[entry.x][entry.y] = '0'
	queue := []pos{entry}

	for step = 1; len(queue) > 0; step++ {
		tmp := queue
		queue = nil

		for _, p := range tmp {
			// 拿到方块后，不判断这个方块，判断周围的方块
			// 因为这个遍历不是DFS，不用判断自身
			for _, d := range delta {
				x, y := p.x+d.x, p.y+d.y
				if 0 <= x && x < m && 0 <= y && y < n && maze[x][y] == '.' {
					// 如果在边界上，返回
					if x == 0 || x == m-1 || y == 0 || y == n-1 {
						return
					}
					maze[x][y] = '0'
					queue = append(queue, pos{x, y})
				}
			}
		}
	}
	return -1
}

// 994. 腐烂的橘子
func orangesRotting(grid [][]int) (step int) {
	type pos struct {
		x, y int
	}
	delta := []pos{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	rows, cols := len(grid), len(grid[0])
	var fresh int

	var queue []pos
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 2 {
				queue = append(queue, pos{i, j})
			} else if grid[i][j] == 1 {
				fresh++
			}
		}
	}

	for step = -1; len(queue) > 0; step++ {
		tmp := queue
		queue = nil

		for _, p := range tmp {
			for _, d := range delta {
				x, y := p.x+d.x, p.y+d.y
				if x >= 0 && x < rows && y >= 0 && y < cols && grid[x][y] == 1 {
					grid[x][y] = 2
					fresh--
					queue = append(queue, pos{x, y})
				}
			}
		}
	}

	if fresh > 0 {
		return -1
	}
	return max(step, 0)
}
