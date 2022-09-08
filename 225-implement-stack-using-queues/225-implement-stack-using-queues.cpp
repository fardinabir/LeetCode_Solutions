class MyStack {
public:
    queue <int> main, aux;
    int last;
    MyStack() {
        
    }
    void push(int x) {
        main.push(x);
        last = x;
    }
    int pop() {
        while(main.size()>1){
            aux.push(main.front());
            main.pop();
        }
        int tmp = main.front();
        main.pop();
        while(!aux.empty()){
            push(aux.front());
            aux.pop();
        }
        return tmp;
    }
    int top() {
        return last;
    }
    
    bool empty() {
        return main.empty();
    }
};
