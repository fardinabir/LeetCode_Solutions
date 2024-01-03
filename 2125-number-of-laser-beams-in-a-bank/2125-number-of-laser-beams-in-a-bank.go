func numberOfBeams(bank []string) int {
    var res, prev int
    for _, row := range bank {
        cnt := countDevice(row)
        if cnt != 0 {
            res += prev * cnt
            prev = cnt
        }
    }
    return res
}

func countDevice(row string) (cnt int) {
    for _, ch := range row {
        if ch == '1' {
            cnt++
        }
    }
    return
}