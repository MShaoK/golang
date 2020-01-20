package main

import "fmt"

func main() {
	defer fmt.Println("world")
	defer fmt.Println("where will this go")
	fmt.Println("hello")
	fmt.Println("4th?")
	defer fmt.Println("first Defer")

}
