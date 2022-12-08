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
	fmt.Printf("Введите входные данные: ")

	reader := bufio.NewReader(os.Stdin)
	a, _ := reader.ReadString('\n')

	// чистим пробелы в строке
	a = strings.TrimSuffix(a, "\n")
	a = strings.ReplaceAll(a, " ", "") // как вариант еще нашел интересное решение
	
	// проверка на недопустимые символы
	matched, _ := regexp.MatchString(`[^0-9,^\+,^\-,^\*,^\/,I,V,X,\ ]`, a)

	if matched {
		fmt.Println("Строка не является математической операцией. Работа приложения завершена.")
		os.Exit(1)
	}

	// проверка на двойной операнд
	if strings.Count(a, "+") > 1 || strings.Count(a, "-") > 1 || strings.Count(a, "*") > 1 || strings.Count(a, "/") > 1 {
		fmt.Println("Формат математической операции не удовлетворяет заданию — два операнда и один оператор. Работа приложения завершена.")
		os.Exit(1)
	}

	fmt.Println("Все в порядке. " + a)

	reg := regexp.MustCompile(`[\+,\-,\*,\/]`)

	digitSplit := reg.Split(a, -1)

	// цикл по символам в строке
	for _, value := range digitSplit {
		fmt.Println(value)
	}

	os.Exit(1)

}
