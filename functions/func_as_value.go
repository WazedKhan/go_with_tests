package main

import "fmt"

// first-class citizen function
func add(x, y int) (sum int) {
	sum = x + y
	return
}

// first-class citizen function
func mul(x, y int) (mul int) {
	mul = x * y
	return
}

// Higher-Order function
func aggregate(a, b, c int, arithmetic func(int, int) int) (secondResult int) {
	firstResult := arithmetic(a, b)
	secondResult = arithmetic(firstResult, c)
	return
}


func main() {
	sum := aggregate(2, 3, 4, add)
	fmt.Println("Sum: ", sum)

	product := aggregate(2, 3, 4, mul)
	fmt.Println("Product: ", product)
}