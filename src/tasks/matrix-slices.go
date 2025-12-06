package tasks

import (
	"fmt"
	"strings"
)

type UserMx struct {
	Name string
}

func m10() {
	var mx [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Matrix 10=", mx)
}

func stringMatrix() {
	var mx [2]string
	mx[0] = "Hello"
	mx[1] = "World"
	fmt.Println("String Matrix = ", mx)
}

func userMx() {
	mx := [2]UserMx{{"Jhon"}, {"Pedro"}}
	fmt.Println("User Matrix = ", mx)
}

func mxSlice() {
	mx := [3]UserMx{{"Jhon"}, {"Pedro"}, {"Andre"}}
	fmt.Println("Slice Matrix Before = ", mx)
	fmt.Println("Slice Matrix After = ", mx[1:])
	fmt.Println("Slice Matrix Test = ", mx[:3])
}

func slices() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("Matrix no literal", q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println("Matrix no literal", r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println("Matrix no literal", s)
}

// ! Diminui, mas se tiver capacidade tu pode fazer um slice aumentando
func sliceCapacity() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)

}

func printSlice(s []int) {
	fmt.Printf("Matrix slice capcity len=%d cap=%d %v\n", len(s), cap(s), s)
}

func emptySlice() {
	var s []int
	fmt.Println("EMPTY SLICE")
	printSlice(s)
	if s == nil {
		fmt.Println("nil!")
	}
}

func makeSlices() {
	a := make([]int, 5)
	fmt.Println("MAKE")
	printSlice(a)
	b := make([]int, 5, 20)
	fmt.Println("MAKE2")
	printSlice(b)
	printSlice(b[:cap(b)])
}

func slicesOfSlices() {
	// Create a tic-tac-toe board.
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"
	board[2][0] = "X"
	board[2][1] = "O"
	board[1][1] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func addingItemToSlice() {
	a := make([]int, 1)
	b := make([]int, 10)
	fmt.Println("APPEND")
	printSlice(a)
	a = append(a, 1, 2, 3)
	printSlice(a)
	fmt.Println("APPEND Ellipses by b", b)
	a = append(a, b...)
	printSlice(a)
}

func loopInSlice() {
	fmt.Println("Loop in Slice")
	a := make([]int, 10)
	for i, v := range a {
		fmt.Println(i, v)
	}
	fmt.Println("Loop in Slice without value")
	for i := range a {
		fmt.Println(i)
	}
	fmt.Println("Loop in Slice without index")
	for _, v := range a {
		fmt.Println(v)
	}
}

func pic(dx, dy int) [][]uint8 {
	x := make([][]uint8, dy)

	for i := range x {
		row := make([]uint8, dx)
		for t := range row {
			row[t] = uint8(t ^ 10*i)
		}
		x[i] = row
	}

	return x
}

func TestMatrix() {
	m10()
	stringMatrix()
	userMx()
	mxSlice()

	// slices
	slices()
	sliceCapacity()
	emptySlice()
	makeSlices()
	slicesOfSlices()
	addingItemToSlice()
	loopInSlice()
	pic(10, 20)
}
