package main

import "fmt"

func max(arr []int, length int) int {
	for i := 0; i < length-1; i++ {
		if arr[i] > arr[i+1] {
			temp := arr[i]
			arr[i] = arr[i+1]
			arr[i+1] = temp
		}
	}

	maxValue := arr[length-1]
	return maxValue
}
func main() {
	scores := []int{12, 23, 34, 24, 6534, 64, 54}
	length := len(scores)
	maxValue := max(scores, length)
	fmt.Println(maxValue)
}
