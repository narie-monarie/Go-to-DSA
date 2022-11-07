class Solution {
    public boolean canJump(int[] nums) {
        int reachable = 0;
        int len = nums.length;
        for(int i = 0; i < nums.length && i <= reachable; i++){
            reachable = Math.max(reachable, i + nums[i]);
            if(reachable >= len - 1){
                return true;
            }
        }
        return false;
    }
}