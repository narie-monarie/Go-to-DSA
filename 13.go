/** Write a Go Program to Find factorial of a Number using For loop.
The for loop (for i := 1; i <= factorialnum; i++) iteration starts at
one and ends at user given value. Within the loop
(factorial = factorial * i), we multiply each value
with a factorial variable value **/

package main

import "fmt"

func factorial(factorialnum int) int {
	max := 1
	for i := 1; i <= factorialnum; i++ {
		max = max * i
	}

	return max
}

func main() {

	var factorialnum, max int
	fmt.Scanf("%d", &factorialnum)
	max = factorial(factorialnum)
	fmt.Printf("The factorial is: %d", max)

}
