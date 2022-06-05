func finalValueAfterOperations(operations []string) int {
    total := 0
    for _, i := range operations{
        if i == "X++" || i == "++X"{
            total++
        }else if i == "--X" || i == "X--"{
            total--
        }
    }
    
    return total
}