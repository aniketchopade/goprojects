package main

import "fmt"

func main() {
	a()
}

func a() {
	fmt.Println("inside a")
	//recover only happens in defer sequence
	//call to recover inside a deferred function stops the panicking sequence
	//restoring normal execution and retrieves the error value passed to the call of panic
	defer func() {
		d := recover()
		if d != nil {
			fmt.Println("Recover inside a()!")
			c()
		}
	}()
	b()
	c()
}

func b() {
	fmt.Println("inside b")
	panic("panic in b")
}
func c() {
	panic("panic in c")
	fmt.Println("inside c")
}
