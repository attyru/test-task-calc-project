package main

import (
	"fmt"
	"regexp"
	"strings"

	//"strconv"
	// "math"
	"bufio"
	"os"
)

func main() {

	//var valuestring string
	//var res float64
	var matched bool
	
	fmt.Printf("Введите входные данные: ")

	reader := bufio.NewReader(os.Stdin)
	a, _ := reader.ReadString('\n')

	// чистим пробелы в строке
	a = strings.TrimSuffix(a, "\n")
	a = strings.ReplaceAll(a, " ", "")

	// проверка на длину строки (минимум 3)
	if len(a) < 3 {
		fmt.Println("Введены некорректные данные. Работа приложения завершена.")
		os.Exit(1)	
	}
	
	// проверка на недопустимые символы
	matched, _ = regexp.MatchString(`[^0-9,^\+,^\-,^\*,^\/,^I,^V,^X]`, a)

	if matched {
		fmt.Println("Строка не является математической операцией. Работа приложения завершена.")
		os.Exit(1)
	}

	// проверка на двойной операнд. кривенькое условие - от одного да 4х раз вызывается функция, но лучше не придумал
	if strings.Count(a, "+") > 1 || strings.Count(a, "-") > 1 || strings.Count(a, "*") > 1 || strings.Count(a, "/") > 1 {
		fmt.Println("Формат математической операции не удовлетворяет заданию — два операнда и один оператор. Работа приложения завершена.")
		os.Exit(1)
	}

	// проверка на то что операнд не является открывающим или замыкающим выражение
	lastsymbol := string(a[len(a)-1])
	firstsymbol := string(a[0])

	if firstsymbol == "0" || firstsymbol == "+" || firstsymbol == "-" || firstsymbol == "*" || firstsymbol == "/" {
		fmt.Println("Первый символ не должен быть операндом или 0. Работа приложения завершена.")
		os.Exit(1)
	}

	if lastsymbol == "+" || lastsymbol == "-" || lastsymbol == "*" || lastsymbol == "/" {
		fmt.Println("Последний символ не должен быть операндом. Работа приложения завершена.")
		os.Exit(1)
	}

	// проверка на то что отсутствует мешанина из арабских и римских цифр
	roman, _ := regexp.MatchString(`[I,V,X]`, a)
	arabic, _ := regexp.MatchString(`[0-9]`, a)

	if roman && arabic {
		fmt.Println("Выражение должно состоять только из арабских или только из римских чисел, разделенных операндами. Работа приложения завершена.")
		os.Exit(1)
	}

	fmt.Println("Все в порядке. " + a)

	// цикл по символам в строке
	// for _, value := range a {
	// 	valuestring = string(value)
	// 	matched, _ := regexp.MatchString(`[^0-9]`, valuestring)
	//digit, err := strconv.Atoi(valuestring)
	//	if matched {
	// тут либо операнд, либо римская цифра, ничего лучше не придумал чем как в свитч затолкать
	// switch {
	// case string(valuestring) == "+":
	// 	if valuestringold == valuestring {

	// 	}

	// }
	// 	continue
	// }

	//valuestringold = valuestring

	//fmt.Printf("%d \n", digit)
	//}

	os.Exit(1)

}
