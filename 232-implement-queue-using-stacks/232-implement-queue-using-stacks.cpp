class MyQueue {
public:
    stack <int> main, aux;
    void push(int x) {
        main.push(x);
    }
    int pop() {
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
    int peek(){
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