package main

import "fmt"

func main() {
	var a, b = 1, 2
	fmt.Printf("%d\n", a+b)
	fmt.Println("aa")
	fmt.Println("aaaaa")
	for index := 0; index < 10; index++ {
		fmt.Println(index)
	}
}

func add(a int, b int) int {
	return a + b
}
