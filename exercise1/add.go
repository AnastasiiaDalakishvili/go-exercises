package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	result := Add(1209, 234562, 7658493)

	fmt.Println(formatNumber(result))
}

func Add(numbers ...interface{}) int {
	sum := 0
	for _, num := range numbers {
		if i, ok := num.(int); ok {
			sum += i
		}
	}
	return sum
}

func formatNumber(num int) string {
	if num > 9999 {
		str := fmt.Sprintf("%d", num)
		re := regexp.MustCompile("(\\d+)(\\d{3})")

		for n := ""; n != str; {
			n = str
			str = re.ReplaceAllString(str, "$1,$2")
		}
		return str
	}
	return strconv.Itoa(num)
}
