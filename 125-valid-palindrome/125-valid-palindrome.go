func isPalindrome(s string) bool {    
    for start, end := 0, len(s)-1; start < end; start, end = start+1,end-1 {
        for start < end && !isAlphaNumeric(int(s[start])) {
            start++
        }
        
        for start < end && !isAlphaNumeric(int(s[end])) {
            end--
        }
        
        if strings.ToLower(string(s[start])) != strings.ToLower(string(s[end])) {
            return false
        }
    }
    
    return true
}

func isAlphaNumeric(asciCode int) bool {
	return (int('a') <= asciCode && asciCode <= int('z')) ||
		(int('A') <= asciCode && asciCode <= int('Z')) ||
		(int('0') <= asciCode && asciCode <= int('9'))
}