package go
func postorderTraversal(root *TreeNode) []int {
    result := []int{}
    
    if root == nil { return []int{} }
    
    result = append(result, postorderTraversal(root.Left)...)
    result = append(result, postorderTraversal(root.Right)...)
    result = append(result, root.Val)
    return result
}

