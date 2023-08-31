package main

import "fmt"

func MakingSlices() {
	a := make([]int, 0, 5)

	a = a[:cap(a)]
	a = a[1:cap(a)]
	a = a[:cap(a)]

	fmt.Println(len(a), cap(a), a)
}
