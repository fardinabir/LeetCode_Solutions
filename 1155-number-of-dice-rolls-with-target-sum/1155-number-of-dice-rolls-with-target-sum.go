var mem [31][1001]int

const mod=1000000007

func numRollsToTarget(n int, k int, target int) int {
    memset2D(&mem, -1)
    return calPossibleWays(n, k, target)
}

func memset2D(arr *[31][1001]int, value int) {
	for i := range arr {
		for j := range arr[i] {
			arr[i][j] = value
		}
	}
}

func calPossibleWays(n, k, target int) int {
    if n == 0 {
        if target == 0 {
            return 1
        }
        return 0
    }
    if target <= 0 {
        return 0
    }
    
    if mem[n][target] != -1 {
        return mem[n][target]
    }
    
    var ways int
    for i := 1; i <= k; i++ {
        // fmt.Println("calling : ", n-1, "---", target - i)
        ways = (ways + (calPossibleWays(n-1, k, target - i) % mod)) % mod
        // fmt.Println(ways)
    }
    mem[n][target] = ways % mod
    return ways
}