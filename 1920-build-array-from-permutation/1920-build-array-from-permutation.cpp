class Solution {
public:
    vector<int> buildArray(vector<int>& nums) {
      vector<int> b;
        for(int y: nums) b.push_back(nums[y]);
        return b;
    }
};