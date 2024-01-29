type MyQueue []int

func Constructor() MyQueue {
    return MyQueue{}
}


func (this *MyQueue) Push(x int)  {
    *this = append(*this, x)
}


func (this *MyQueue) Pop() int {
    front := (*this)[0]
    *this = (*this)[1:]
    return front
}


func (this *MyQueue) Peek() int {
    return (*this)[0]
}


func (this *MyQueue) Empty() bool {
    return len((*this)) == 0
}