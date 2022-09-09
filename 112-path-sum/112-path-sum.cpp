class Solution {
public:
    bool hasPathSum(TreeNode* root, int targetSum) {
        if(!root || (!root->left && !root->right)) return root? targetSum==root->val : false;
        return hasPathSum(root->left, targetSum - root->val)|hasPathSum(root->right, targetSum - root->val);
    }
};