func isPalindrome(x int) bool {
    var rev, rem int
    for i := x; i > 0; i /= 10 {
        rem = i % 10
        rev = rev*10 + rem
    }
    
    if x == rev{
        return true
    }
    
    return false
}