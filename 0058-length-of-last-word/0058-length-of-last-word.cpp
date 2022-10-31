class Solution {
public:
    int lengthOfLastWord(string s) {
        stringstream ss(s);
        string word;
        vector<string>x;
        while(ss>>word){
            x.push_back(word);
        }
        
        return x.back().size();
    }
};