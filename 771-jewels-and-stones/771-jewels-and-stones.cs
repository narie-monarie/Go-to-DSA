public class Solution {
    public int NumJewelsInStones(string jewels, string stones) {
        return stones.Count(c => jewels.Contains(c));
    }
}