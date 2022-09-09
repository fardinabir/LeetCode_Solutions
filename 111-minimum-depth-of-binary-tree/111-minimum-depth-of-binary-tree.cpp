class Solution {
public:
    int minDepth(TreeNode* root) {
        if(root==NULL)
            return 0;
        int a = 10e6, b = 10e6;
        if(root->left)
            a = minDepth(root->left);
        if(root->right)
            b = minDepth(root->right);
        if(a==10e6 && b==10e6)
            return 1;
        else
            return min(a,b)+1;
    }
};