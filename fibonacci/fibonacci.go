package main

import "fmt"

func main() {
	value1 := int64(1)
	value2 := int64(2)
	sum := int64(2)

	for {
		result := value1 + value2
		if result >= 4000000 {
			break
		}
		if result%2 == 0 {
			sum += result
		}

		value1 = value2
		value2 = result
	}
	fmt.Printf("sum = %v\n", sum)
}
