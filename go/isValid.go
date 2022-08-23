package main

func isValidBST(root *TreeNode) bool {
    return isValid(root, math.MinInt64, math.MaxInt64)
}

func isValid(node *TreeNode, min, max int) bool {
	if node == nil {
		return true
	}
	return node.Val > min && node.Val < max &&
		isValid(node.Left, min, node.Val) &&
		isValid(node.Right, node.Val, max)
}