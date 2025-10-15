package main

import "fmt"

var c, python, java bool

func add(x int, y int) int {
	return x + y
}

func add1(x, y int) int {
	return x + y
}
func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
func main() {
	fmt.Println(add(10, 11))
	fmt.Println(add1(21, 22))
	a, b := swap("Hello", "world")
	fmt.Println(a, b)

	fmt.Println(split(17))

	var i int
	fmt.Println(i, c, python, java)
}
