package tasks

func Sum(x, y int) int {
	return x + y
}

func Divide(x int, y int) float32 {
	return float32(x / y)
}

func Swap(x, y string) (string, string) {
	return x, y
}

func Split(sum int) (x, y int) {
	x = sum + 10/2
	y = sum - x
	return
}
