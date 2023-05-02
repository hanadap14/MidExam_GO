// ANSWER //
// In Go language programming, the ellipsis (aka "three dots" or "spread operator") is used 
// to indicate that a function can take 
// a variable number of arguments of a particular type. 
// This is known as variadic arguments. 
// The syntax for using an ellipsis in Go is to place it before the type of the 
// variadic parameter in the function definition. 
// For example : 

package main

import "fmt"

func average(numbers ...float64) float64 {
	sum := 0.0
	count := 0.0

	for _, num := range numbers {
		sum += num
		count++
	}

	if count > 0 {
		return sum / count
	} else {
		return 0.0
	}
}

func main() {
	fmt.Println(average(1.5, 2.0, 2.5)) 
	fmt.Println(average(5.5, 6.5)) 
	fmt.Println(average()) 
}
