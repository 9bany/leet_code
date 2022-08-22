package main

func preorderTraversal(root *TreeNode) []int {
    if root == nil {
		return nil
	}
	stack := []*TreeNode{root}
	result := make([]int, 0)
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, cur.Val)
		if cur.Right != nil {
			stack = append(stack, cur.Right)
		}
		if cur.Left != nil {
			stack = append(stack, cur.Left)
		}
	}
	return result
}