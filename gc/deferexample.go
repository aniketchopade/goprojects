package main

import "fmt"

//D1 function
func D1() {
	for i := 3; i > 0; i-- {
		defer fmt.Print(i, " ")
	}
}

//D2 Function
func D2() {
	for i := 3; i > 0; i-- {
		defer func() {
			fmt.Print(i, " ")
		}()
	}
}

//D3 function
func D3() {
	for i := 3; i > 0; i-- {
		defer func(n int) {
			fmt.Print(n, " ")
		}(i)
	}
}
