package tasks

import "fmt"

func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}

type List[T any] struct {
	next *List[T]
	val  T
}

func TestGenerics() {
	fmt.Println(Index([]int{10, 20, 15, -10}, 100))
	fmt.Println(Index([]string{"Pedro"}, "Pedro"))
}
