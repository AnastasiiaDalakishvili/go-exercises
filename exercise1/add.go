package main

import "fmt"

func main() {
	result := Add(5, 4, 3, 2, -10)
	fmt.Println(result)
}

func Add(numbers ...int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}
