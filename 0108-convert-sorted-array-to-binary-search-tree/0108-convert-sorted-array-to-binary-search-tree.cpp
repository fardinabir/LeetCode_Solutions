class Solution {
public:
    TreeNode* makeTree(int l, int r, vector<int> &nums) {
        if(l>r){
            return NULL;
        }
        int mid = (l+r)/2;
        TreeNode* node = new TreeNode(nums[mid]);
        
        node->left = makeTree(l, mid-1, nums);
        node->right = makeTree(mid+1, r, nums);
        
        return node;
    }
    
    TreeNode* sortedArrayToBST(vector<int>& nums) {
        
        return makeTree(0, nums.size()-1, nums);
    }
};