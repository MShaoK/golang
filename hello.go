package main

import (
	"fmt"
	"math/rand"
	"time"
	"math"

)

func doNothing(x, y, z int) (int, int, int){
	return x , y , z
}

func nakedReturn(sum int) (x, y int){
	x = sum + 10
	y = sum - 10
	return

}

func main() {
	fmt.Println(nakedReturn(10))
	fmt.Println(doNothing(50, 100, 1))
	fmt.Println("Welcome to the playground")
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	fmt.Println(math.Pi)
	fmt.Println("The time is", time.Now())
}

