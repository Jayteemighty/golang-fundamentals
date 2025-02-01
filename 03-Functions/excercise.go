// Try writing a function that takes a list of numbers and returns both the sum and the average of those numbers.

package main

import "fmt"

func sumAndAverage(numbers []float64) (sum float64, average float64) {
	for _, value := range numbers {
		sum += value
	}
	if len(numbers) > 0 {
		average = sum / float64(len(numbers))
	}
	return
}

func main() {
	numbers := []float64{12.0, 23.0, 34.0, 54.0, 56.0, 76.0}
	sum, average := sumAndAverage(numbers)
	fmt.Println("the sum = ", sum)
	fmt.Println("the average = ", average)
}
