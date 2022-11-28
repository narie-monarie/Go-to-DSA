class Solution {
    public int majorityElement(int[] arr) {
        
    		if (arr == null || arr.length == 0)
			return Integer.MIN_VALUE;

		int count = 1, maj = arr[0];

		for (int i = 0; i < arr.length; i++) {
			if (arr[i] == maj) {
				count++;
			} else {
				count--;
			}

			if (count == 0) {
				maj = arr[i];
				count = 1;
			}
		}

		return maj;
    }
}