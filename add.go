package main

import (
	"flag"
	"fmt"
	"regexp"
	"strconv"

	"github.com/AnastasiiaDalakishvili/go-exercise/adapters"
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
		numbersFromFile := adapters.GetNumberFromFile(filePath)
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

//func csvFileReader(filePath string) []int {
//	numbersFromFile := []int{}
//
//	f, err := os.Open(filePath)
//
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	r := csv.NewReader(f)
//
//	for {
//		record, err := r.Read()
//		if err == io.EOF {
//			break
//		}
//		if err != nil {
//			log.Fatal(err)
//		}
//
//		for _, i := range record {
//			j, err := strconv.Atoi(i)
//			if err != nil {
//				panic(err)
//			}
//			numbersFromFile = append(numbersFromFile, j)
//		}
//
//		return numbersFromFile
//	}
//	return numbersFromFile
//}
