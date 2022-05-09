func twoSum(nums []int, target int) []int {
    m := make(map[int]int)
    
    for i, j := range nums{
        if j, ok := m[target - j]; ok{
            return []int{j, i}
        }
        m[j] = i
    }
    return []int{}
}