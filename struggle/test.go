package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sufficientSubset1(root *TreeNode, limit int) *TreeNode {
	dummy := &TreeNode{
		Val:  0,
		Left: root,
	}
	execute(root, dummy, 0, limit, 0)
	return dummy.Left
}

func execute(root *TreeNode, father *TreeNode, direction int, limit int, sum int) bool { //返回值是root点是否被删
	if nil == root {
		return false
	}
	sum += root.Val
	left := execute(root.Left, root, 0, limit, sum)
	right := execute(root.Right, root, 1, limit, sum)
	// 如果是叶子节点并且路径和小于limit 或 左子树被删并且没有右子树 或 右子树被删并且没有左子树 或左右子树都被删 则删掉root
	if (nil == root.Left && nil == root.Right && sum < limit) ||
		(left && nil == root.Right) || (right && nil == root.Left) || (left && right) {
		if 0 == direction {
			father.Left = nil
		} else {
			father.Right = nil
		}
		return true
	}
	return false
}
