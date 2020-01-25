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
	array := [5]int {1,2,3,4,5}
	fmt.Println(array)
	slice := array[2:]
	fmt.Println(slice, "slice of array?")
	fmt.Println(cap(slice), "capacity ")
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}

