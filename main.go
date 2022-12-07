package main

import (
	"fmt"
	"regexp"
	"strings"
	// "math"
	"bufio"
	"os"
)

func main() {

	//var res float64
	fmt.Print("Введите входные данные: ")

	reader := bufio.NewReader(os.Stdin)
	a, _ := reader.ReadString('\n')
	a = strings.TrimSuffix(a, "\n")
  
	a = strings.ReplaceAll(a, " ", "")
	  
	// проверка на недопустимые символы
	matched, _ := regexp.MatchString(`[^0-9,^\+,^\-,^\*,^\/,I,V,X,\ ]`, a)
  
	if matched {
	  fmt.Println("Строка не является математической операцией. Работа приложения завершена.")
	  os.Exit(1)
	}
  
	fmt.Println("Все в порядке. " + a)

	// 	switch op {

	// 	case "!":
	// 		res = factorial(int(a))

	// 	case "^":
	// 		res = math.Pow(a, b)

	// 	case "+":
	// 		res = a + b
	// 	case "-":
	// 		res = a - b
	// 	case "*":
	// 		res = a * b
	// 	case "/":
	// 		if b == 0 {
	// 			fmt.Println("Делить на ноль нельзя!")
	// 			os.Exit(1)
	// 		}
	// 		res = a / b

	// 	default:
	// 		fmt.Println("Операция выбранна не верно")
	// 		os.Exit(1)
	// 	}

	// 	fmt.Printf("Результата выполнения операции: %.2f\n", res)
	// }

	// // функция вывода факториала
	// func factorial(n int) float64 {
	// 	var res = 1
	// 	for i := 1; i <= n; i++ {
	// 		res *= i
	// 	}
	// 	return float64(res)
}
