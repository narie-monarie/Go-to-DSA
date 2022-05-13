package main

import "fmt"

func main() {
	names := make([]string, 3, 6)
	//length -> 3
	//capacity -> 6

	names[0] = "Coat"
	names[1] = "LifeJacket"
	names[2] = "Paddle"

	fmt.Println("len: ", len(names))
	fmt.Println("len: ", cap(names))
}
