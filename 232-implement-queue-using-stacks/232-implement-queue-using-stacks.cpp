class MyQueue {
public:
    stack <int> main, aux;
    MyQueue() {
        
    }
    
    void push(int x) {
        main.push(x);
    }
    
    int pop() {
        while(!aux.empty())
            aux.pop();
        while(!main.empty()){
            aux.push(main.top());
            main.pop();
        }
        int top = aux.top();
        aux.pop();
        while(!aux.empty()){
            main.push(aux.top());
            aux.pop();
        }
        return top;
    }
    int peek() {
        while(!aux.empty())
            aux.pop();
        while(!main.empty()){
            aux.push(main.top());
            main.pop();
        }
        int top = aux.top();
        while(!aux.empty()){
            main.push(aux.top());
            aux.pop();
        }
        return top;
    }
    
    bool empty() {
        return main.empty();
    }
};

/**
 * Your MyQueue object will be instantiated and called as such:
 * MyQueue* obj = new MyQueue();
 * obj->push(x);
 * int param_2 = obj->pop();
 * int param_3 = obj->peek();
 * bool param_4 = obj->empty();
 */