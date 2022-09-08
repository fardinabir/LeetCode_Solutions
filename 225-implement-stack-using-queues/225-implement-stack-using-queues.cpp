class MyStack {
public:
    queue <int> main, aux;
    int last;
    MyStack() {
        
    }
    void push(int x) {
        int sz = main.size();
        main.push(x);
        while(sz--){
            main.push(main.front());
            main.pop();
        }
    }
    int pop() {
        int tmp = main.front();
        main.pop();
        return tmp;
    }
    int top() {
        return main.front();
    }
    
    bool empty() {
        return main.empty();
    }
};
