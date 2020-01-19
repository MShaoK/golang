package main

import "fmt"

const (
	Big = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

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
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
}

