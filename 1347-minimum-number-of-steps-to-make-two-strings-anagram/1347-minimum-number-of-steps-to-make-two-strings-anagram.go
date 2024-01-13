func minSteps(s string, t string) (res int) {
    mapChar := [26]int{}
    for k , v := range s {
        mapChar[v - 'a']++
        mapChar[t[k] - 'a']--
    }

    for _ , v := range mapChar {
        if v > 0 { 
            res += v
        }
    }
    return res
}