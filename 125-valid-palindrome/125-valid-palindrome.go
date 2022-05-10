func isPalindrome(s string) bool {
  bs := []byte(s)
	n := 0
	for _, b := range bs {
		if ('a' <= b && b <= 'z') ||
			('A' <= b && b <= 'Z') ||
			('0' <= b && b <= '9') {
			bs[n] = b
			n++
		}
	}

	bsl := strings.ToLower(string(bs[:n]))

	for i, j := 0, len(bsl)-1; i <= j; i, j = i+1, j-1 {
		if bsl[i] != bsl[j] {
			return false
		}
	}

	return true
}

//or

func isPalindrome(s string) bool {
	s = strings.ToLower(s)

	r, _ := regexp.Compile("[^a-z0-9]+")

	s = r.ReplaceAllString(s, "")

	start, end := 0, len(s)-1

	for start < end {
		if s[start] != s[end] {
			return false
		}

		start++
		end--
	}

	return true
}
