package main

func invertTree(root *TreeNode) *TreeNode {
    invert(root)
    return root
}

func invert(root *TreeNode) {
    if root != nil {
        tmp := root.Left
        root.Left = root.Right
        root.Right = tmp

        invert(root.Left)
        invert(root.Right)
    }
}