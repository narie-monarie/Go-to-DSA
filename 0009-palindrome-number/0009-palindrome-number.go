func isPalindrome(n int) bool {
    	if n < 0 || (n%10 == 0 && n != 0) {
		return false
	}
	reverse := 0

	for n > reverse {
		pop := n % 10
		reverse = reverse*10 + pop
		n /= 10
	}
	return n == reverse || n == reverse/10
}