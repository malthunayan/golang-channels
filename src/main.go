package main

import (
	"fmt"
)

// This is a slice, with each element denoting an arbitrary number of rows. You can
// change the numbers in the slice to test the program's output, and it should always
// yield (rowNumber)^3, for each element in rowNumbers. Also, note that this normally
// wouldn't have package scope, but it's here for easier access, to test the code.
var rowNumbers = []int{5, 1, 7, 3}

func main() {
	// A channel is made
	ch := make(chan int)

	// This is the function that calculates the rowNumber (an element in the
	// rowNumbers slice) cubed using for loops, and its input is a sending channel.
	go rowSumOddNumbers(ch)

	// This is the function that prints out the result to the terminal,
	// and its input is a receiving channel.
	printResult(ch)
}

// rowSumOddNumbers sums the odd numbers in a given row for the odd
// triangle, so row three of the triangle would consist of: 9, 11, 13
// and their sum would be 9+11+13 = 27 = 3^3
func rowSumOddNumbers(ch chan<- int) {
	for _, rowNumber := range rowNumbers {

		// The following two for loops are a lengthy way of calculating
		// rowNumber^3, it is mainly just to show that I know how to use for
		// loops, and it is how the sum is calculated in the odd triangle.
		index := 0
		for i := 1; i < rowNumber; i++ {
			index += i
		}
		initial := 2*index + 1
		add := initial
		acc := initial
		for i := 1; i < rowNumber; i++ {
			add += 2
			acc += add
		}

		// To make the program run much faster comment out the previous
		// two for loops, and uncomment the following line:
		// acc = rowNumber * rowNumber * rowNumber

		// After calculating the row's total, it is sent to the channel.
		ch <- acc
	}

	// The channel is closed and this signals the for loop in printResult
	// that there are no more channels being sent.
	close(ch)
}

// Prints the resulting summations to the terminal
func printResult(ch <-chan int) {
	// This method ranges over the number of channels and prints out
	// each sum, or (element)^3 for each element in rowNumbers.
	for sum := range ch {
		fmt.Println(sum)
	}
}
