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