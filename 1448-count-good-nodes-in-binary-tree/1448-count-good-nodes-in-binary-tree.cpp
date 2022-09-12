class Solution {
public:
    int goodNodes(TreeNode* root) {
        return findRes(root,root->val);
    }
    int findRes(TreeNode *node, int mx){
        if(!node) return 0;
        int downRes = findRes(node->left, max(mx,node->val)) + findRes(node->right, max(mx,node->val));
        return downRes + (mx <= node->val);
    }
};