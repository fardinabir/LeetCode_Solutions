func findMatrix(nums []int) [][]int {
    var freq [201]int
    
    var matrix [][]int
    for _, num := range nums {
        freq[num]++
        
        if freq[num] > len(matrix) {
            newRow := []int{num}
            matrix = append(matrix, newRow)
        } else {
            posArray := &matrix[freq[num]-1]
            *posArray = append(*posArray, num)
        }
    }
    
    fmt.Println(matrix)
    return matrix
}