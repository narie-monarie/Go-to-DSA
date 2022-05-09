func isSubsequence(s string, t string) bool {
    left := 0
    left2 := 0
    
    for left < len(s) && left2 < len(t){
        if s[left] == t[left2]{
            left++
        }
        left2++
    }
            if left == len(s){
            return true
        }
    return false
}