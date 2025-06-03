package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// flattenWithTail 返回当前子树展开后的最后一个节点
func flattenWithTail(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	// 递归 flatten 左右子树
	leftTail := flattenWithTail(root.Left)
	rightTail := flattenWithTail(root.Right)

	// 如果左子树存在，把它插入到右边，然后接上原来的右子树
	if root.Left != nil {
		leftTail.Right = root.Right
		root.Right = root.Left
		root.Left = nil
	}

	// 返回最后的尾节点（右子树尾巴优先）
	if rightTail != nil {
		return rightTail
	}
	if leftTail != nil {
		return leftTail
	}
	return root // 如果左右都没有，自己是叶子节点
}

// flatten 是对外的接口，封装内部递归逻辑
func flatten(root *TreeNode) {
	flattenWithTail(root)
}

// 打印链表
func printList(root *TreeNode) {
	for root != nil {
		fmt.Printf("%d ", root.Val)
		root = root.Right
	}
	fmt.Println()
}

// 示例
func main() {
	root := &TreeNode{1,
		&TreeNode{2,
			&TreeNode{3, nil, nil},
			&TreeNode{4, nil, nil}},
		&TreeNode{5,
			nil,
			&TreeNode{6, nil, nil}},
	}

	flatten(root)
	printList(root) // 输出: 1 2 3 4 5 6
}
