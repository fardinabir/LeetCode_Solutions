class Solution {
public:
    vector<vector<int>> levelOrder(TreeNode* root) {
        vector <vector <int> > vc;
        if(root==NULL)
            return vc;
        queue <TreeNode *> tq;
        tq.push(root);
        while(!tq.empty()){
            vector <int> add;
            int lsz = tq.size();
            for(int i=0; i<lsz; i++){
                TreeNode *temp = tq.front();
                tq.pop();
                add.push_back(temp->val);
                if(temp->left!=NULL)
                    tq.push(temp->left);
                if(temp->right!=NULL)
                    tq.push(temp->right);
            }
            vc.push_back(add);
        }
        return vc;
    }
};