
package main
import "fmt"

func main() {
	result := fibarray(5)
	for ix, fib := range result {
		fmt.Printf("The %d-th Fibonacci number is: %d\n", ix,  fib)
	}
}

func fibarray(term int) []int {	// calculating Fibonacci for first term numbers
	farr := make([]int, term)

	//base case
	if term == 1 {
		farr[0] = 0
		return farr
	}
	farr[0], farr[1] = 0, 1

	// recursive case
	for i:= 2; i < term; i++ {
		farr[i] = farr[i-1] + farr[i-2]	// calculating sequence for i index
	}
	return farr
}
