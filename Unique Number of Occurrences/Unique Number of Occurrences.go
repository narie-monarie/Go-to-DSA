package main

func uniqueOccurrences(arr []int) bool {
	count := make(map[int]int, len(arr)) //make a hashtable same length as the arr and index them
	for _, i := range arr {
		count[i]++ // add the items in the hash map
	}
	m := make(map[int]bool, len(count)) //go thru the hashtable this is a set and indicating the duplicates
	for _, j := range count {
		if m[j] {
			return false
		}
		m[j] = true
	}
	return true
	i
}
