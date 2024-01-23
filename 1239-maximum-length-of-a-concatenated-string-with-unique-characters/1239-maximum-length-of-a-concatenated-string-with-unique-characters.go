func maxLength(arr []string) int {
    return findComb(0, arr, "")
}

func findComb(index int, arr []string, st string) int {
    if index == len(arr) {
        return getUniqueCharCount(st)
    }
    taken := findComb(index+1, arr, st + arr[index])
    notTaken := findComb(index+1, arr, st)
    return max(taken, notTaken)
}

func getUniqueCharCount(st string) int {
    var charMap [26]int
    for _, x := range st {
        charMap[x-'a'] ++
        if charMap[x-'a'] > 1 {
            return 0
        }
    }
    var count int
    for _, val := range charMap {
            count += val
    }
    return count
}