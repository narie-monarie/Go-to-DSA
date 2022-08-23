public class Solution {
    public int RemoveDuplicates(int[] nums) {
        if(nums.Length == 0) return 0;
        int a = 0;
        
        for(int i = 1; i < nums.Length; i++){
            if(nums[a] != nums[i]){
                a++;
                nums[a] = nums[i];
            }
        }
        return a + 1;
    }
}