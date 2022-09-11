class Solution {
public:
    bool isValidBST(TreeNode* root) {
        return evaluate(root, LONG_MAX, LONG_MIN);
    }
    bool evaluate(TreeNode *node, long mx, long mn){
        if(!node)
            return true;
        if(node->val <= mn || node->val >= mx)
            return false;
        return evaluate(node->left, node->val, mn)&&evaluate(node->right, mx, node->val);
    }
};