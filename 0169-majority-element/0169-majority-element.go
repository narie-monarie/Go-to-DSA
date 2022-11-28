func majorityElement(nums []int) int {
    if len(nums) == 0 {
        return 0;
    }    
    
    count := 0
    maj := nums[0]
    
    for i := 0; i < len(nums); i++ {
        if nums[i] == maj {
            count++
        }else{
            count--
        }
        
        if count == 0 {
            maj = nums[i]
            count = 1
        }
    }
    
    return maj
}