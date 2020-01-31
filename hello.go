package main

import (
	"fmt"
)

type Test struct {
	firstname string
	lastname string
	age int
	assets float64
}

func main() {
	colin := Test{"person", "test", 55, 123.23}
	testMap := make(map[string]int)
	testMap["Colin"] = 10

	fmt.Println(colin)
	fmt.Println(testMap)
	

}

