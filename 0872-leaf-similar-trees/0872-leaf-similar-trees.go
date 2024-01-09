/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
    return findLeaf(root1, 0) == findLeaf(root2, 0) 
}

func findLeaf(root *TreeNode, base int) int {
    if root == nil {
        return base
    }
    if root.Left == nil && root.Right == nil {
        return (base+root.Val) * 10
    }
    return findLeaf(root.Right, findLeaf(root.Left, base))
}