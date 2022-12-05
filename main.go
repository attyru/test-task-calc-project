package main

import (
	"fmt"
	"math"
	"os"
)

func main() {

	var a, b, res float64
	var op string

	fmt.Print("Введите число: ")
	fmt.Scanln(&a)

	fmt.Print("Введите арифметическую операцию (+, -, *, /, ^, !: ")
	fmt.Scanln(&op)

	fmt.Print("Введите второе число: ")
	fmt.Scanln(&b)

	switch op {

	case "!":
		res = factorial(int(a))

	case "^":
		res = math.Pow(a, b)

	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		if b == 0 {
			fmt.Println("Делить на ноль нельзя!")
			os.Exit(1)
		}
		res = a / b

	default:
		fmt.Println("Операция выбранна не верно")
		os.Exit(1)
	}

	fmt.Printf("Результата выполнения операции: %.2f\n", res)
}

// функция вывода факториала
func factorial(n int) float64 {
	var res = 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	return float64(res)
}
