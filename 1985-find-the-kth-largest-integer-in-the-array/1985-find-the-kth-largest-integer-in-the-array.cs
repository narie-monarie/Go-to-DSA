public class Solution {
    public string KthLargestNumber(string[] nums, int k) {
        return nums.OrderBy(A => A.Length).ThenBy(A => A).ElementAt(nums.Length-k);
    }
}