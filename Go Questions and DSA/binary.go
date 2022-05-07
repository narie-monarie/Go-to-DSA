package main

func binarySearch(arr[]int,target){
	left , right := 0, len(arr)-1
	for(left<=right){
		mid := (left + right)%2

		if (arr[mid] == target){
			return mid
		}else if(arr[mid] < target){
			left = mid + 1
		}else{
			right = mid - 1
		}
	}
}