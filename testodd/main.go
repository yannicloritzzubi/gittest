package main

import "fmt"

func createSlice() []int {

	var s []int
	for i := 0; i < 11; i++ {
		s = append(s, i)
	}

	return s
}

func main() {
	s := createSlice()

	fmt.Println(s)

	for num := range s {
		if num%2 == 0 {
			fmt.Println(num, " is even")
		} else {
			fmt.Println(num, " is odd")
		}
	}
}
