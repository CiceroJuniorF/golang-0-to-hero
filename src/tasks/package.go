package tasks

import (
	"fmt"
	"math"
	"math/rand"
)

func FavoriteNumberRand() {
	fmt.Println("Your favorite number is:", rand.Intn(200))
}

func YourProblems() {
	fmt.Printf("Now you have: %g problems\n", math.Sqrt(7))
}
