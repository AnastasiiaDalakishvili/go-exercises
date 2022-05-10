package adapters

import (
	"regexp"
	"strconv"
)

func DataConverter(input string) []int {
	var str []string
	numbersFromFile := []int{}
	a := regexp.MustCompile(`(\s*(,|\n)\s*)`)
	str = a.Split(input, -1)

	//range through the numbers and append then to array of int
	for _, i := range str {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		numbersFromFile = append(numbersFromFile, j)
	}

	return numbersFromFile
}
