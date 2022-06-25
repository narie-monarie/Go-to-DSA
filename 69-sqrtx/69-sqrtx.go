import "math"
func mySqrt(x int) int {
    a := math.Sqrt(float64(x))
    v := math.Floor(a)
    return int(v)
}