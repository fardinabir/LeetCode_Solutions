func halvesAreAlike(s string) bool {
    return checkVowels(s[:len(s)/2]) == checkVowels(s[(len(s)/2):])
}

func checkVowels(str string) (cnt int) {
    for _, x := range str {
        if strings.ContainsRune("aeiouAEIOU", x) {
            cnt++
        }
    }
    return
}