package tasks

import "fmt"

func AwaitTheFuncEnd() {
	defer fmt.Println("World!")
	fmt.Print("Hello ")
}

func StackOfDefer() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
