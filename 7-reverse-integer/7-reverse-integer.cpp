class Solution {
public:
    int reverse(int x) {
        bool positive=true;
        if(x<0){
            if(x!=-2147483648)
            x=-x;
            else
                return 0;
            positive = false;
        }
        long long curr=0;
        while(x>0){
            int rem = x%10;
            curr = curr*10 + rem;
            x/=10;
        }
        if(!positive)
            curr = -curr;
        if(curr  > 2147483647 || curr < -2147483648){
             return 0;   
        }
        int vl = curr;
        return vl;
    }
};