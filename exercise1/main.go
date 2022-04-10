package main

import "fmt"

func main() {
	result := Add(5, 4)
	fmt.Println(result)
}

func Add(n, y int) int {
	return n + y
}
