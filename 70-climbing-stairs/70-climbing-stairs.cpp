class Solution {
public:
    int climbStairs(int n) {
        if(n<=2)
           return n;
        int cls[n];
        cls[0] = 1;
        cls[1] = 2;
        for(int i=2; i<n; i++)
            cls[i] = cls[i-1] + cls[i-2];
        return cls[n-1];
    }
};