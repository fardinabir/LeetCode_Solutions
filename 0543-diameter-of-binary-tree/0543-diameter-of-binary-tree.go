func diameterOfBinaryTree(root *TreeNode) int {
    _, mdm := findMaxDepth(root)
    return mdm
}

func findMaxDepth(node *TreeNode) (mDepth, mDiameter int) {
    if node == nil {
        return 0, 0
    }
    l, ldm := findMaxDepth(node.Left)
    r, rdm := findMaxDepth(node.Right)
    mDepth = max(l, r) + 1
    mDiameter = max(l + r, max(ldm, rdm))
    return mDepth, mDiameter
}