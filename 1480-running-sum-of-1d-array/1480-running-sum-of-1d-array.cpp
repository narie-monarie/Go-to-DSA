class Solution {
public:
    vector<int> runningSum(vector<int>& nums) {
        vector<int>z;
        int a = 0;
        for(int y :nums){
            a+=y;
            z.push_back(a);
        }
        return z;
    }
};