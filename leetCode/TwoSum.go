package main

import "fmt"

/*func twoSum(arr []int, target int) []int {
	for i := 0; i < len(arr)-1; i++ {
		for j := i; j < len(arr); j++ {
			if arr[i]+arr[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}

}*/

//sorted array
func twoSum(arr []int, target int) []int {
	if len(arr) <= 1 {
		return []int{}
	}

	var left, right int = 0, len(arr) - 1

	for left < right {
		sum := arr[right] + arr[left]
		if sum == target {
			return []int{left, right}
		} else if sum > target {
			right--
		} else {
			left++
		}
	}

	return []int{}

}

func main() {
	scores := []int{1, 8, 9, 11, 12, 23}
	fmt.Println(twoSum(scores, 32))
}
