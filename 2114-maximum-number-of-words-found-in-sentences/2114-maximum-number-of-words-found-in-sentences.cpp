class Solution {
public:
    int mostWordsFound(vector<string>& sentences) {
        int n = 0;
        for(int i = 0; i < sentences.size(); i++){
            int c_max = 1;
            for(int j = 0; j < sentences[i].size(); j++){
                if(sentences[i][j] == ' '){
                    c_max++;
                }
            }
            n = max(n , c_max);
        }
        return n;
    }
};


// class Solution {
// public:
//     int mostWordsFound(vector<string>& sentences) {
//         int cur_max = 0;
//         for(int i = 0; i < sentences.size(); i++){
//             int count = 1;
//             for(int j = 0; j < sentences[i].size(); j++){
//                 if(sentences[i][j] == ' '){
//                     count++;
//                 }
//             }
//             cur_max = max(cur_max, count);
//         } 
//         return cur_max;   
//     }
// };
