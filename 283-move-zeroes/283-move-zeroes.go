func moveZeroes(nums []int)  {
    
    left := 0
    
    for i := 0; i < len(nums); i++{
        if nums[i] != 0{
            if i != 0{
                nums[i], nums[left] = nums[left], nums[i]
                left++
            }else{
                left++
            }
        }
    }
    
}