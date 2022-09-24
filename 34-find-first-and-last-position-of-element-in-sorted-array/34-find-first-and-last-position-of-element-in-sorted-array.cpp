class Solution {
public:
    vector<int> searchRange(vector<int>& nums, int target) {
        int i = -1, j = -1;
        for(int k=0;k<nums.size();k++){
            if(i==-1 && nums[k]==target)
                i = k;
            if(nums[k]==target)
                j = k;
            if(i!=-1 && nums[k]!=target) break;
        }
        return {i, j};
    }
};