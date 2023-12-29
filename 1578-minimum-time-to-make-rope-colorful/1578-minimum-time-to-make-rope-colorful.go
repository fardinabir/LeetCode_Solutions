func minCost(colors string, neededTime []int) int {
    var cal, j int
    for i:= 0; i < len(colors); {
        maxTime := neededTime[i]
        cum := neededTime[i]
        for j = i+1; j < len(colors) && colors[j] == colors[i]; j++ {
            // fmt.Println(neededTime[j],"------", neededTime[i])
            maxTime = max(maxTime, neededTime[j])
            cum += neededTime[j]
            // fmt.Println("min : ", maxTime)
        }
        cum -= maxTime
        i = j
        cal += cum
        // fmt.Println("-=-=-=", cum)
    }
    
    return cal
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}