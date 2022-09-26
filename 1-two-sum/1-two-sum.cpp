class Solution {
public:
    vector<int> twoSum(vector<int>& nums, int target) {
        map<int,int> m;
        int k;
        for(int i = 0; i < nums.size(); i++){
            k = target - nums[i];
            if (m.find(k) != m.end()){
                return {i,m[target - nums[i]]};
            }
            m[nums[i]] = i;
        }
        return {};
    }
};