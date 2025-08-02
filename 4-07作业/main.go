package main

import (
	"fmt"
	"math"
	"sort"
	"time"
)

type sumType interface {
	sum() float64
}

type graphic struct {
	long int
}

type rotundity struct { // 圆形
	graphic
	pai float64
}

func (r *rotundity) sum() float64 {
	return r.pai * float64(r.long*r.long)
}

type rectangle struct { // 长方形
	graphic
	wide int
}

func (r *rectangle) sum() float64 {
	if r.long == r.wide {
		panic("长方形的长宽不一致")
	} else {
		return float64(r.long * r.wide)
	}
}

type square struct { // 正方形
	graphic
	wide int
}

func (s *square) sum() float64 {
	return float64(s.long * s.wide)
}

func typeSum() {
	typeSlice := []sumType{
		&rotundity{graphic: graphic{long: 3}, pai: math.Pi},
		&rectangle{graphic: graphic{long: 3}, wide: 4},
		&square{graphic: graphic{long: 3}, wide: 3},
	}
	sumList := make([]float64, 0, len(typeSlice))
	for i, _ := range typeSlice {
		var inter sumType = typeSlice[i]
		fmt.Println(inter.sum())
		sumList = append(sumList, inter.sum())
	}
	sort.Slice(sumList, func(i, j int) bool {
		return sumList[i] > sumList[j]
	})
	fmt.Println(sumList)
}

var funcMap = make(map[int]int, 10)

func fib(n int) {
	switch {
	case n < 0:
		fmt.Println("input int is negtive")
	case n == 0:
		funcMap[0] = 0
	case n == 1:
		funcMap[1] = 1
	case n == 2:
		funcMap[2] = 1
	case n >= 3:
		funcMap[n] = funcMap[n-1] + funcMap[n-2]
	}
}

func tUnix() {
	beforeTime := "2017/08/02 09:30:00 +0800"
	t, err := time.Parse("2006/01/02 15:04:05 -0700", beforeTime)
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("1、毫秒为%v\n", t.UnixMilli())
	}
	t1 := t.Add(12 * time.Hour)
	s1 := t1.Format("2006/01/02 15:04:05 -0700")
	fmt.Printf("2、格式化时间%v\n", s1)
	res1 := t.Weekday()
	res2 := time.Since(t1)
	fmt.Printf("3、周几: %v 过了多少周：%v\n", res1, int(res2.Hours()/24/7))
	fmt.Printf("4、到今天过了多少天: %v\n", int(res2.Hours()/24))
}

func main() {
	typeSum()
	fmt.Println("--------------")
	testLevel := 45
	for i := 0; i < testLevel; i++ {
		fib(i)
		fmt.Printf("%d ", funcMap[i])
	}
	fmt.Println("--------------")
	tUnix()
}
