func maxAncestorDiff(root *TreeNode) int {
    return findMaxDiff(root, root.Val, root.Val)
}

func findMaxDiff(node *TreeNode, mx, mn int) int {
    if node == nil {
        return mx - mn
    }
    mx = max(mx, node.Val)
    mn = min(mn, node.Val)
    return max(findMaxDiff(node.Left, mx, mn),
                         findMaxDiff(node.Right, mx, mn))
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
