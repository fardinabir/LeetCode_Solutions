func pseudoPalindromicPaths (root *TreeNode) int {
    var curState [10]int
    return traverseTree(root, curState)
}

func traverseTree(node *TreeNode, curState [10]int) int {
    if node == nil {
        return 0;
    }
    curState[node.Val]++
    if node.Left == nil && node.Right == nil {
        return isPseudoPalindrome(&curState);
    }
    return traverseTree(node.Left, curState) + traverseTree(node.Right, curState)
}

func isPseudoPalindrome(curState *[10]int) int {
    var odd int
    for _, x := range curState {
        if x%2 == 1 {    
            odd++
        }
    }
    if odd > 1 {
        return 0
    }
    return 1
}