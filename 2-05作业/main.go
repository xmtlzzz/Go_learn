package main

import (
	"fmt"
	"math/rand"
)

func sliceAppend() {
	s1 := make([]int, 3, 4)
	s2 := append(s1, 1)
	fmt.Printf("%v %d\n", s1, len(s1))
	fmt.Printf("%v %d\n", s2, len(s2))
	s1[1] = 2
	fmt.Printf("%v\n", s1)
	fmt.Printf("%v\n", s2)
	s2[3] = 2
	fmt.Printf("%v\n", s1)
	fmt.Printf("%v\n", s2)
	fmt.Println("-------")
	fmt.Printf("s1: %p len: %d cap: %d\n", &s1[0], len(s1), cap(s1))
	fmt.Printf("s2: %p len: %d cap: %d\n", &s2[0], len(s2), cap(s2))
	s2 = append(s2, 3)
	fmt.Printf("s1: %p len: %d cap: %d\n", &s1[0], len(s1), cap(s1))
	fmt.Printf("s2: %p len: %d cap: %d\n", &s2[0], len(s2), cap(s2))
}

func arrayAddtoSlice() {
	a := [...]int{1, 4, 9, 16, 2, 5, 10, 15}
	b := make([]int, len(a)-1)
	for i := 0; i < len(a)-1; i++ {
		b[i] = a[i] + a[i+1]
	}
	fmt.Println(b)
}

func IntSort() {
	a := make([]int, 0, 100)
	b := make(map[int]int, len(a))
	for i := 0; i < 100; i++ {
		ranInt := rand.Intn(200) - 100
		a = append(a, ranInt)
		if _, ok := b[ranInt]; !ok {
			b[ranInt] = 1
		} else {
			b[ranInt] += 1
		}
	}
}

func main() {
	// sliceAppend()
	// arrayAddtoSlice()
	IntSort()
}
