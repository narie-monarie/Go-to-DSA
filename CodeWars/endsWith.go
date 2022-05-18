package main

import (
	"fmt"
	"strings"
)

func solution(str, ending string) bool {
	if strings.Contains(str, ending) {
		return true
	}
	return false
}

func main() {
	fmt.Println(solution("abc", "d"))
}
