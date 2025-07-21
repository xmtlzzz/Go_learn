package main

import (
	"fmt"
	"strconv"
	"strings"
)

var builder strings.Builder

func markDG(n int) (res int) {
	switch n {
	case 0:
		return 1
	case 1:
		return 1
	case 2:
		return n * (n - 1)
	}
	res = markDG(n - 1)
	return n * res
}

func formatIntList(n int) {
	res := ""
	maxlen := 1
	for i := n; i <= n; i++ {
		for j := i; j >= 1; j-- {
			builder.WriteString(strconv.Itoa(j))
			if j > 1 {
				builder.WriteString(" ")
			}
		}
		a := builder.String()
		maxlen = len(a)
		builder.Reset()
	}
	for i := 1; i <= n; i++ {
		for j := i; j >= 1; j-- {
			builder.WriteString(strconv.Itoa(j))
			if j > 1 {
				builder.WriteString(" ")
			}
		}
		res = builder.String()
		fmt.Printf("%*s\n", maxlen, res)
		builder.Reset()
	}

}

func main() {
	// fmt.Println(markDG(5))
	formatIntList(13)
}
