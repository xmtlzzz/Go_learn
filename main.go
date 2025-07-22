package main

import "fmt"

func fibloop(n int) int {
	switch {
	case n < 0:
		fmt.Println("input int is negtive")
	case n == 0:
		return 0
	case n == 1 || n == 2:
		return 1
	}
	a, b := 1, 1
	for i := 0; i < n-2; i++ {
		a, b = b, a+b
	}
	return b
}

func fib(n int) int {
	switch {
	case n < 0:
		fmt.Println("input int is negtive")
	case n == 0:
		return 0
	case n == 1 || n == 2:
		return 1
	}
	return fib(n-1) + fib(n-2)
}

func fibDG(n, a, b int) int {
	switch {
	case n < 0:
		fmt.Println("input int is negtive")
	case n == 0:
		return 0
	case n == 1 || n == 2:
		return b
	}
	return fibDG(n-1, b, a+b)

}

func main() {
	for i := 0; i < 10; i++ {
		// fmt.Println(fibloop(i))
		// fmt.Println(fib(i))
		fmt.Println(fibDG(i, 1, 1))
	}
}
