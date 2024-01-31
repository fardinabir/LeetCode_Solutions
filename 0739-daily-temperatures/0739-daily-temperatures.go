func dailyTemperatures(temperatures []int) []int {
	var stackArr []int
	waitingDays := make([]int, len(temperatures), len(temperatures))

	for i := 0; i < len(temperatures); i++ {
		for len(stackArr) > 0 && temperatures[stackArr[len(stackArr)-1]] < temperatures[i] {
			prevIndex := stackArr[len(stackArr)-1]
			stackArr = stackArr[:len(stackArr)-1]
			waitingDays[prevIndex] = i - prevIndex
		}
		stackArr = append(stackArr, i)
	}

	return waitingDays
}