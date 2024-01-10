/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var graph map[int][]int
var visited map[int]bool

func amountOfTime(root *TreeNode, start int) int {
    graph = make(map[int][]int)
    visited = make(map[int]bool)
    
    prepareGraph(root)
    
    return findMaxDepth(start, 0)
}

func prepareGraph(node *TreeNode) {
    if node == nil {
        return
    }
    
    if node.Left != nil {
        graph[node.Val] = append(graph[node.Val], node.Left.Val)
        graph[node.Left.Val] = append(graph[node.Left.Val], node.Val)
    }
    if node.Right != nil {
        graph[node.Val] = append(graph[node.Val], node.Right.Val)
        graph[node.Right.Val] = append(graph[node.Right.Val], node.Val)
    }
    prepareGraph(node.Left)
    prepareGraph(node.Right)
}

func findMaxDepth(curNode,depth int) int {
    if _, ok := visited[curNode]; ok {
        return depth-1
    }
    visited[curNode] = true
    
    mx := 0
    for _, x := range graph[curNode] {
        mx = max(mx, findMaxDepth(x, depth+1))
    }
    
    return mx
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

