/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode() : val(0), left(nullptr), right(nullptr) {}
 *     TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
 *     TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
 * };
 */
class Solution {
public:
    TreeNode* insertIntoBST(TreeNode* root, int val) {
        if(root == NULL) return new TreeNode(val);;
        if(root->val > val && insertIntoBST(root->left, val)->val == val)
            root->left = new TreeNode(val);
        else if(root->val < val && insertIntoBST(root->right, val)->val == val)
            root->right = new TreeNode(val);
        return root;
    }
};