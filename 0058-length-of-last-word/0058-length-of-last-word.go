func lengthOfLastWord(s string) int {
    a := strings.Trim(s," ")
    b := strings.Split(a, " ")
    c := b[len(b)-1]
    return len(c)
}