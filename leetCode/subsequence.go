package main

import "strings"

// 2 solutions and 1 leetcode solution
//0(n) time 0(1) space
func subseq(arr []int, seq []int) {
	arr_left := 0
	seq_left := 0

	for arr_left < len(arr) && seq_left < len(seq) {
		if arr[arr_left] == seq[seq_left] {
			seq_left++
		}
		arr_left++
	}
	return seq_left == len(seq)
}

//2
func subseq(arr []int, seq []int) {
	seq_left := 0

	for value := range arr {
		if seq_left == len(seq) {
			if seq[seq_left] == value {
				seq_left++
			}
		}
	}
	return seq_left == len(seq)
}

//leetcode subsequence
func isSubsequence(s string, t string) bool {
	first := strings.Split(s, "")
	second := strings.Split(t, "")

	first_left := 0
	second_left := 0
	for first_left < len(first) && second_left < len(second) {
		if first[first_left] == second[second_left] {
			first_left++
		}
		second_left++
	}

	return first_left == len(first)

}
