func makeEqual(words []string) bool {
    var freq [26]int
    for _, str := range words {
        for _, ch := range str {
            freq[ch-'a']++
        }
    }
    
    for _, fr := range freq {
        if fr % len(words) != 0 {
            return false
        }
    }
    return true
}