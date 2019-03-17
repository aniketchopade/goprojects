package main

import "fmt"

func main() {
	a()
}

func a() {
	fmt.Println("inside a")
	defer func() {
		d := recover()
		if d != nil {
			fmt.Println("Recover inside a()!")
			//c()
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
	fmt.Println("inside c")
}
