func findContentChildren(g []int, s []int) int {
    sort.Ints(s)
    sort.Ints(g)
    i:=0
    for j:=0; j< len(s) && i< len(g); j++ {
        if(g[i] <= s[j]){
            i += 1
        }
    }
    return i
}