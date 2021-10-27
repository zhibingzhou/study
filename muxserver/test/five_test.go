package test

import (
	"fmt"
	"testing"
)

func add(arg ...int) {
	var sum int
	for _, value := range arg {
		sum += value
	}
	fmt.Println(sum)
}

func TestUseALL(t *testing.T) {

	//Remove(1)

	abc := []int{1, 2, 3, 4, 5}
	fmt.Println(abc[0:4])
	first := abc[0:5]

	n := 2
	f := abc[0:n]
	e := abc[n+1:]
	mid := []int{}
	mid = append(mid, f...)
	mid = append(mid, e...)

	//end := abc[0:4]

	end := append(abc[0:4], abc[5:]...)

	fmt.Println("first", first)
	fmt.Println("mid", mid)
	fmt.Println("end", end)

	switch 0 {
	case 0:
		fmt.Println("4560")
		fallthrough
	case 1:
		fmt.Println("fallthrough")
	}

	add([]int{123, 123}...)

	add(123, 123)
}

func Remove(value interface{}) {

	abc := []int{1, 2, 3, 4, 5}
	result := []int{}
	for i, v := range abc {
		if value == v {
			result = append(abc[:i], abc[i+1:]...)
		}
	}
	fmt.Println(value, result)
}


type Integer int
