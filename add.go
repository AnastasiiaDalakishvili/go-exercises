package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

func main() {
	// Declare a string flag called file path with a default value ("input.txt") and a help message
	fileName := flag.String("input-file", "", "specify file name")
	flag.Parse()

	path := *fileName

	result := Add(path)
	fmt.Println(formatNumber(result))
}

//Add calculates the sum of the numbers passed to it. If no data passed to the function then it reads it from the file
func Add(filePath string, numbersPassedtoFunction ...interface{}) int {
	sum := 0
	if numbersPassedtoFunction != nil {
		//calculate sum
		for _, num := range numbersPassedtoFunction {
			if i, ok := num.(int); ok {
				sum += i
			}
		}
	} else {
		numbersFromFile := getNumberFromFile(filePath)
		for _, num := range numbersFromFile {
			sum += num
		}
	}
	return sum
}

//this function formats the number to thousands separated by comma
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

//Read the input file and return an array of numbers
func getNumberFromFile(filePath string) []int {
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var perline int
	var nums []int

	for {
		_, err := fmt.Fscanf(file, "%d\n", &perline) // give a patter to scan

		if err != nil {
			if err == io.EOF {
				break // stop reading the file
			}
			fmt.Println(err)
			os.Exit(1)
		}
		nums = append(nums, perline)
	}
	return nums
}