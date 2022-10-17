package main

/**
*   @author wangchenyang
*   @date 2022/9/30 10:26
*   @description:
 */
func sufficientSubset(root *TreeNode, limit int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		// 叶子节点
		if root.Val < limit {
			return nil
		}
		return root
	}
	root.Left = sufficientSubset(root.Left, limit-root.Val)
	root.Right = sufficientSubset(root.Right, limit-root.Val)
	if root.Left != nil || root.Right != nil {
		// 左右子树只要有一颗不为空，本节点就存在，否则不存在
		return root
	}
	return nil
}
