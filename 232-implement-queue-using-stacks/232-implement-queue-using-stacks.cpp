class MyQueue {
public:
    stack <int> main, aux;
    int front;
    void push(int x) {
        if(main.empty())
            front = x;
        main.push(x);
    }
    int pop(){
        if(aux.empty()){
            while(!main.empty()){
                aux.push(main.top());
                main.pop();
            }
        }
        int top = aux.top();
        aux.pop();
        return top;
    }
    int peek(){
        if(!aux.empty())
            return aux.top();
        return front;
    }
    bool empty() {
        return main.empty() && aux.empty();
    }
};