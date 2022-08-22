package go

func maxDepth(root *TreeNode) int {
    if root == nil { return 0 }
    
    lengtLeft := maxDepth(root.Left)
    lengtRight := maxDepth(root.Right)
    
    if lengtLeft > lengtRight {
        return lengtLeft + 1
    } else {
        return lengtRight + 1
    }
}