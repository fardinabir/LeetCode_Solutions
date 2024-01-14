
func closeStrings(word1 string, word2 string) bool {
    mpWord1 := make(map[byte]int)
    mpWord2 := make(map[byte]int)
    if len(word1) != len(word2) {
        return false
    }
    for i:=0; i<len(word1); i++ {
        mpWord1[word1[i]]++
        mpWord2[word2[i]]++
    }
    mpArr1 := []int{}
    mpArr2 := []int{}
    for key, val := range mpWord1 {
        var val2 int
        var ok bool
        if val2, ok = mpWord2[key]; !ok {
            return false
        }
        mpArr1 = append(mpArr1, val)
        mpArr2 = append(mpArr2, val2)
    }
    
    slices.Sort(mpArr1)
    slices.Sort(mpArr2)
    
    if slices.Equal(mpArr1, mpArr2) {
        return true
    }
    return false
}

