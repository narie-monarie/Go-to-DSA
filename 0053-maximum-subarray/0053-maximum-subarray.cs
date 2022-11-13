public class Solution {
    public int MaxSubArray(int[] nums) {
        int ans = int.MinValue;
        int a = 0;
        foreach(int x in nums){
            a += x;
            ans = Math.Max(a, ans);
            a = Math.Max(a, 0);
        }
        return ans;
    }
}