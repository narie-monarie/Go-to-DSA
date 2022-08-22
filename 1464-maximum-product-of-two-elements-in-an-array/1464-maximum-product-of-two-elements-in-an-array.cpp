class Solution {
public:
    int maxProduct(vector<int>& nums) {
        sort(nums.begin(),nums.end());
        return (nums.back()-1)*((nums.rbegin()[1])-1);
    }
};