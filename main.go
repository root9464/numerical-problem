package main

import "fmt"

func void(array []int) []int {
	n := len(array) - 1

	for n >= 0 && array[n] == 9 {
		array[n] = 0
		n--
	}
	if n < 0 {
		return append([]int{1}, array...)
	} else {
		array[n]++
	}
	return array
}

func main() {
	num := []int{1234}
	fmt.Println(void(num))
}
