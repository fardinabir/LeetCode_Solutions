class Solution {
public:
    bool isSymmetric(TreeNode* root) {
        return check(root->left,root->right);
    }
    bool check(TreeNode* l,TreeNode* r){
        if(l==nullptr || r==nullptr)
            return (l==nullptr && r==nullptr);
        return (l->val==r->val)&(check(l->left,r->right))&(check(l->right,r->left));
    }
    
};