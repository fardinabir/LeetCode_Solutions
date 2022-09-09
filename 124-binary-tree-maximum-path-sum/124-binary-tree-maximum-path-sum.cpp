class Solution {
public:
    int maxPathSum(TreeNode* root) {
        pair <int,int> pi; // first is cont. and second is max
        pi = findPath(root);
        return max(pi.first, pi.second);
    }
    pair<int,int> findPath(TreeNode *node) {
        if(node == NULL)
            return {-10e4, -10e4};
        pair <int,int> pLeft,pRight;
        pLeft = findPath(node->left);
        pRight = findPath(node->right);
        
        int cont, mx;
        cont = max(max(pLeft.first + node->val, pRight.first + node->val), node->val);
        mx = max(max(max(pLeft.second, pRight.second),pLeft.first + node->val + pRight.first), cont);
        return {cont, mx};
    }
};