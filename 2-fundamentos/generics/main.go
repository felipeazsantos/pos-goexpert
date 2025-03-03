package main

import "fmt"

type MyNumber int64

type Number interface {
	~int64 | ~float64
}

func sum[T Number](nums ...T) T {
	var total T
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	var num1 MyNumber = 10
	var num2 MyNumber = 5
	s := sum(num1, num2)

	var num3 float64 = 50.53
	var num4 float64 = 4.07
	s2 := sum(num3, num4)

	fmt.Println("The sum value (int64) is:", s)
	fmt.Println("The sum value (float64) is:", s2)
}
