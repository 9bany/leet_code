package main

func levelOrder(root *TreeNode) [][]int { 
    result := [][]int{}
    
    var dns func(*TreeNode, int)
    
    dns = func(root *TreeNode, level int) {
        if root == nil {
            return 
        }
        
        if level >= len(result) {
            result = append(result, []int{})
        }
        
        result[level] = append(result[level], root.Val)
		dns(root.Left, level+1)
		dns(root.Right, level+1)
        
    }
    dns(root, 0)
    return result
}