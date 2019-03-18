package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("you are using", runtime.Compiler)
	fmt.Println("on ..", runtime.GOARCH)
	fmt.Println("running under ", runtime.GOOS)
	fmt.Println("memory allocation recorded after every ", runtime.MemProfileRate)
	fmt.Println("Total logical cpu ", runtime.NumCPU())

}
