class Solution {
public:
    vector<int> inorderTraversal(TreeNode* root) {
        vector <int> vc;
        traverse(root,vc);
        return vc;
    }
    void traverse(TreeNode* node, vector <int> &vc){
        if(node == NULL)
            return;
        traverse(node->left,vc);
        vc.push_back(node->val);
        traverse(node->right,vc);
    }
};