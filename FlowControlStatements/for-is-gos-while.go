package main

import "fmt"

func main() {
	sum := 1
	for sum < 3 {
		sum += sum
	}
	fmt.Println(sum)
}
