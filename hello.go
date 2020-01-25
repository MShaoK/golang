package main

import (
	"fmt"
	"math"
)



type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	slice := []int {2,3,5}
	fmt.Println(len(slice))
	fmt.Println(cap(slice), "capacity ")
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}

