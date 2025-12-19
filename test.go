package main

import "fmt"

func Add(a int, b int) int {
	return a + b
}
func Subtract(a int, b int) int {
	return a - b
}
func Multiply(a int, b int) int {
	return a * b
}
func Divide(a int, b int) int {
	if b == 0 {
		return 0 // simple handling for division by zero
	}
	return a / b
}
func main() {
	fmt.Println("Add(2,3) =", Add(2, 3))
	fmt.Println("Subtract(5,2) =", Subtract(5, 2))
	fmt.Println("Multiply(3,4) =", Multiply(3, 4))
	fmt.Println("Divide(10,2) =", Divide(10, 2))
	Nh := "计算结果："

	fmt.Println(Nh, Add(7, 7), Divide(4, 2))
}
