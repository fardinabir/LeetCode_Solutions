class Solution {
public:
    vector<vector<int>> levelOrder(TreeNode* root) {
        vector <vector <int> > vc;
        traverse(root, vc, 0);
        return vc;
    }
    void traverse(TreeNode *node, vector <vector <int> > &vc, int level){
        if(node == NULL)
            return;
        if(level>=vc.size())
            vc.push_back({});
        vc[level].push_back(node->val);
        traverse(node->left,vc,level+1);
        traverse(node->right,vc,level+1);
    }
};