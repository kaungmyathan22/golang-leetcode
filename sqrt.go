package main
func MySqrt(x int) int {
    if x == 0 || x == 1 {
        return x
    }
    
    start, end := 1, x
    var result int
    
    for start <= end {
        mid := start + (end - start) / 2
        
        // If mid*mid == x, we have found the result
        if mid*mid == x {
            return mid
        }
        
        // Since we need floor value of the sqrt, update result
        if mid*mid < x {
            start = mid + 1
            result = mid
        } else {
            end = mid - 1
        }
    }
    
    return result
}
