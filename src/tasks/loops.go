package tasks

import (
	"fmt"
)

func AllInstructionsFor() {
	sum := 0
	for i := 0; sum <= 100; i++ {
		sum += i * 2
		fmt.Println("LoopAll", "[", i, "]", " Sum: ", sum)
	}
}

func SomeInstructionsFor() {
	sum := 0
	i := 0
	for sum <= 100 {
		i++
		sum += i * 2
		fmt.Println("LoopSome", "[", i, "]", " Sum: ", sum)
	}
}
