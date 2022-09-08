class Solution {
public:
    vector<int> preorderTraversal(TreeNode* root) {
        vector <int> vc;
        traverse(root,vc);
        return vc;
    }
    void traverse(TreeNode* node, vector <int> &vc){
        if(node == NULL)
            return;
        vc.push_back(node->val);
        traverse(node->left,vc);
        traverse(node->right,vc);
    }
};