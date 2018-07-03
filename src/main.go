package main

import "fmt"

func main() {
	var i1, i2, i3 int
	var s1 string
	fmt.Scan(&i1)
	fmt.Scan(&i2, &i3)
	fmt.Scan(&s1)

	fmt.Printf("%d %s", i1+i2+i3, s1)
}
