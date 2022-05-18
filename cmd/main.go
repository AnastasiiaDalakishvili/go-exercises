package main

import (
	"github.com/AnastasiiaDalakishvili/go-exercise/adapters"
	"github.com/AnastasiiaDalakishvili/go-exercise/domain"
)

func main() {
	//extract the numbers pass by user either from the file or just an array of integers
	numbers := adapters.GetNumberFromFile()

	//pass this numbers ro calculator and get the sum
	sum := domain.Add(numbers)

	//pass sum to formatter and get desired formatted sum
	formattedSum := domain.FormatNumber(sum)

	//print sum
	adapters.Printer(formattedSum)
}
