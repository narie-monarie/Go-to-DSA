public class Solution {
    public string KthLargestNumber(string[] nums, int k) {
        return nums.OrderBy(x => x.Length).ThenBy(x => x).ElementAt(nums.Length-k);
    }
}