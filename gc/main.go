package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var mem runtime.MemStats
	printStats(mem)

	for i := 0; i < 15; i++ {
		s := make([]byte, 10000000)
		if s == nil {
			fmt.Println("memory allocation failed")
		}
		printStats(mem)
		time.Sleep(5 * time.Second)
	}

	printStats(mem)

}
func printStats(mem runtime.MemStats) {
	runtime.ReadMemStats(&mem)
	fmt.Println("mem.Alloc:", mem.Alloc)
	fmt.Println("mem.TotalAlloc:", mem.TotalAlloc)
	fmt.Println("mem.HeapAlloc:", mem.HeapAlloc)
	fmt.Println("mem.NumGC:", mem.NumGC)
	fmt.Println("-----")
}
