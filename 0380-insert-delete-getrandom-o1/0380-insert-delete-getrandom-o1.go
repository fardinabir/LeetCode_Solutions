type RandomizedSet struct {
    MapSet map[int]bool
}


func Constructor() RandomizedSet {
    return RandomizedSet{MapSet : make(map[int]bool)}
}


func (this *RandomizedSet) Insert(val int) bool {
    if _, ok := this.MapSet[val]; ok {
        return false
    }
    this.MapSet[val] = true
    return true
}


func (this *RandomizedSet) Remove(val int) bool {
    if _, ok := this.MapSet[val]; ok {
        delete(this.MapSet, val)
        return true
    }
    return false
}


func (this *RandomizedSet) GetRandom() int {
    randomIndex := rand.Intn(len(this.MapSet))
    i := 0
    for key := range this.MapSet {
        if i == randomIndex {
            return key
        }
        i++
    }
    return -1
}


/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */