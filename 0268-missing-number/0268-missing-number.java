class Solution {
    public int missingNumber(int[] nums) {
        int n = (nums.length-1) + 1;
        int x = (nums.length-1) - n;
        int m = (nums.length-1) - x;
        int sum = 0;
        for(int i = 0 ;i < nums.length; i++){
            sum+=nums[i];
        }
        
        int v = (m * ( m + 1) / 2) - sum;
        
        return v;
    }
}