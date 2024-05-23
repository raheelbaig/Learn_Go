package main

import "fmt"

func sub(x int, y int) int {
	return x - y
}

func main() {
	fmt.Println(sub(4, 5))

	fmt.Println(sum(10, 20))

}

func sum(x, y int) int {
	return x + y
}
