package main

import "fmt"

func isUnique(arr []int) bool {
	m := make(map[int]bool)
	for _, j := range arr {
		if _, value := m[j]; !value {
			m[j] = true
		}
	}
	if len(arr) == len(m) {
		return true
	}

	return false
}

func main() {
	arr := []int{2, 5, 3, 3, 4, 9, 7}
	fmt.Println(isUnique(arr))
}
