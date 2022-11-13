public class Solution {
    public int MaxSubArray(int[] nums) {
        int maxSub = nums[0];
        int sum = 0;
        for(int i = 0; i < nums.Length; i++){
            if(sum < 0 )
                sum = 0;
            sum += nums[i];
            maxSub = Math.Max(maxSub, sum);
        }
        
        return maxSub;
    }
}