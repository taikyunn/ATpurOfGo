package main

import "fmt"

type Vertex struct {
	x int
	y int
}

func main() {
	v := Vertex{1, 2}
	p := &v
	p.x = 19
	fmt.Println(v)
}
