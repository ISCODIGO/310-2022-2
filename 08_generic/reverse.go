package main

import "fmt"

func main() {
	var arr []int = []int{1, 2, 3}
	var rev_arr []int = Reverse(arr)
	fmt.Println(rev_arr)
}

func Reverse(input []int) []int {
    var output []int

    for i := len(input) - 1; i >= 0; i-- {
        output = append(output, input[i])
    }

    return output
}

func GenericReverse[T any](input []T) []T {
	var output []T
	
	for i := len(input) - 1; i >= 0; i-- {
		output = append(output, input[i])
	}
	return output
}