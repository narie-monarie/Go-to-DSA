func sortedSquares(nums []int) []int {
    l := 0
    r := len(nums)-1
    index := r
    m := make([]int, len(nums))
    
    for l <= r{
        a := nums[l] * nums[l]
        b := nums[r] * nums[r]
        if a > b{
            m[index] = a
            index--
            l++
        }else{
            m[index] = b
            index--
            r--
        }
    }
    return m
}