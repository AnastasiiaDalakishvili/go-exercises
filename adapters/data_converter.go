package adapters

import (
	"fmt"
	"regexp"
	"strconv"
)

func DataConverter(input string) []int {
	// we get an input which is a string with numbers, commas and new lines
	fmt.Printf("input %q\n", input)

	//variables
	var arrayOfnumbers []string
	var numbersFromFile []int

	//get rid of commas and new lines to get array of strings
	a := regexp.MustCompile(`(\s*(,|\n)\s*)`)
	arrayOfnumbers = a.Split(input, -1)
	fmt.Printf("array of strings that are numbers %q\n", arrayOfnumbers)

	for _, number := range arrayOfnumbers {
		integer, _ := strconv.Atoi(number)
		numbersFromFile = append(numbersFromFile, integer)
	}

	return numbersFromFile
}
