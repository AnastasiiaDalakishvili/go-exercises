package main

import (
	"fmt"
)

func main() {
	result := Add(1, 2, 3, 4.2)
	fmt.Println(result)
}

func Add(numbers ...interface{}) interface{} {
	sum := 0
	for _, num := range numbers {
		if i, ok := num.(int); ok {
			sum += i
		}
	}
	return sum
}
