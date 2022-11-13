class Solution {
    public int maxSubArray(int[] nums) {
        int min = Integer.MIN_VALUE;
        int z = 0;
        for(int i : nums){
            z+=i;
            min = Math.max(min,z);
            z = Math.max(z,0);
        }
        return min;
    }
}