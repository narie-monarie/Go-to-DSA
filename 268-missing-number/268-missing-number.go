func missingNumber(nums []int) int {
    n := len(nums)
    var sum int
    for i:=0;i<n;i++{
        sum+=nums[i]
    }
    
    x := (n * (n+1)/2) - sum
    
    return x
}