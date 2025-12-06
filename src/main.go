package main

import (
	"fmt"
	"strconv"

	"github.com/ccoder/golang-0-to-hero/src/tasks"
)

func swap(a, b string) string {
	x, y := tasks.Swap(a, b)
	return x + " " + y
}

func split(number int) string {
	x, y := tasks.Split(number)
	return strconv.Itoa(x) + " " + strconv.Itoa(y)
}

func main() {
	fmt.Println("Starting...")
	tasks.Hello()
	tasks.Time()
	tasks.FavoriteNumberRand()
	tasks.YourProblems()
	fmt.Println("The sum 1+2 is:", tasks.Sum(1, 2))
	fmt.Println("The division between 2/2 is:", tasks.Divide(2, 2))
	fmt.Println("Swap:", swap("2", "2"))
	fmt.Println("Split:", split(100))
	fmt.Println(tasks.GlobalAndLocalVars())
	tasks.EmptyInitialize()
	tasks.AllInstructionsFor()
	tasks.SomeInstructionsFor()
	fmt.Println("Conditional sqrt", tasks.ConditionalSqrt(100), tasks.ConditionalSqrt(-100))
	fmt.Println("Short Declaration Pow", tasks.ShortDeclarationPow(3, 2, 10), tasks.ShortDeclarationPow(3, 3, 20))
	fmt.Println("Short Declaration Pow With Else", tasks.ShortDeclarationPowWithElseAlert(3, 2, 10), tasks.ShortDeclarationPowWithElseAlert(3, 3, 20))
	tasks.SwitchCase("print")
	tasks.SwitchCase("any")
	tasks.WhenIsSaturday()
	tasks.Greetings()
	tasks.AwaitTheFuncEnd()
	tasks.StackOfDefer()
	p := 10
	fmt.Println("Pointers - Value Before Mutiply to 2 by reference", p)
	tasks.MultiplyPointer(&p, 2)
	fmt.Println("Pointers - Value After Mutiply to 2 by reference", p)
	tasks.TestStruct()
	tasks.TestMatrix()
	tasks.MakeASimpleMap()
	tasks.FuncValue()
	tasks.ClosureFunction()
	tasks.Fibonacci()
	tasks.MethodsExec()
	tasks.ExecInterfaces()
	tasks.RunError()
	tasks.TestGenerics()

	//tasks.internal() -> Don't works because it's a private function
}
