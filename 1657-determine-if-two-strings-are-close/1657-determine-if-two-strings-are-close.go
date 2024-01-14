func closeStrings(word1 string, word2 string) bool {
    if len(word1) != len(word2) {
        return false
    }
    
    f1 := getFreq(word1)
    f2 := getFreq(word2)

    for i := 0; i < 26; i++ {
        if f1[i] * f2[i] == 0 && f1[i] + f2[i] > 0{
            return false
        }
    }
    
    slices.Sort(f1)
    slices.Sort(f2)
    
    if slices.Equal(f1, f2) {
        return true
    }
    
    return false
    
}

func getFreq(word string) []int {
    f := make([]int, 26)
    for _, c := range word {
        f[c - 'a']++
    }
    
    return f
}

