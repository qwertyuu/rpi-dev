package main

import (
	"fmt"
)

var (
	buildTime string
)

func main() {
	fmt.Printf("Build Time: %s\n", buildTime)
	sum := 0
	for i := 1; i < 1000; i++ {
		if i%5 == 0 || i%3 == 0 {
			sum += i
		}
	}
	fmt.Println(sum)
}
