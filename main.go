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

	// чистим пробелы в строке
	a = strings.TrimSuffix(a, "\n")
	a = strings.ReplaceAll(a, " ", "")

	// проверка на недопустимые символы
	matched, _ := regexp.MatchString(`[^0-9,^\+,^\-,^\*,^\/,I,V,X,\ ]`, a)

	if matched {
		fmt.Println("Строка не является математической операцией. Работа приложения завершена.")
		os.Exit(1)
	}

	fmt.Println("Все в порядке. " + a)

}
