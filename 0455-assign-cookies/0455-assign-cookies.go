func findContentChildren(g []int, s []int) int {
    sort.Ints(g)
    sort.Ints(s)
    
    i, j, cnt := 0, 0, 0
    
    for ; i < len(g) && j < len(s); j++ {
        if g[i] <= s[j] {
            cnt++
            i++
        }
    }
    
    return cnt
}