class Solution {
public:
    bool isPalindrome(string s) {
        vector<char>y;
        for(int i = 0; i < s.length(); i++){
            if(iswalnum(s[i])){
                y.push_back(tolower(s[i]));
            }
        }
        string m = string(y.begin(), y.end());
        string z = string(y.rbegin(), y.rend());
        if(m == z){
            return true;
        }
        return false;
    }
};