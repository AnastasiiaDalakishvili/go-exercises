package main

import (
	"github.com/AnastasiiaDalakishvili/go-exercise/adapters"
	"github.com/AnastasiiaDalakishvili/go-exercise/domain"
)

func main() {

	sum := domain.Add()
	formattedSum := domain.FormatNumber(sum)
	adapters.Printer(formattedSum)
}
