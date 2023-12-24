func isPathCrossing(path string) bool {
    vertical, horizontal := 0, 0
    location := make(map[string]bool)
    location["0,0"] = true
    for _, s := range path {
        if s == 'N' {
            vertical--
        } else if s == 'S' {
            vertical++
        } else if s == 'E' {
            horizontal++
        } else {
            horizontal--
        }
        key := strconv.Itoa(vertical) + "," + strconv.Itoa(horizontal)
        if location[key] {
            return true
        }
        location[key] = true
    }
    return false
}