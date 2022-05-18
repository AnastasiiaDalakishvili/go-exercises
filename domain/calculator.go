package domain

func Add(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

//sum := 0
//if numbersPassedtoFunction != nil {
////calculate sum
//for _, num := range numbersPassedtoFunction {
//if i, ok := num.(int); ok {
//sum += i
//}
//}
//} else {
//numbersFromFile := adapters.GetNumberFromFile()
//for _, num := range numbersFromFile {
//sum += num
//}
//}
//return sum
