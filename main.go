package main

import "fmt"

func main() {
	a := map[int]int{
		1: 1,
		2: 1,
		3: 2,
		4: 3,
		5: 3,
	}
	n := 0
	for v := range a {
		n += 1
		fmt.Printf("set的第%d元素为: %d\n", n, v)
	}
}
