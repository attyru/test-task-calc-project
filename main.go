package main

import (
	"fmt"
	"regexp"
	"strings"
	"strconv"
	"bufio"
	"os"
	romannumeral "github.com/brandenc40/romannumeral"
)

func main() {
	
	fmt.Printf("Введите входные данные: ")

	reader := bufio.NewReader(os.Stdin)
	a, _ := reader.ReadString('\n')

	// чистим пробелы в строке
	a = strings.TrimSuffix(a, "\n")
	a = strings.ReplaceAll(a, " ", "")

	// проверка на длину строки (минимум 3, максимум 5)
	lengt := len(a)

	if lengt < 3 || lengt > 9 {
		fmt.Println("Введены некорректные данные. Работа приложения завершена.")
		os.Exit(1)	
	}
	
	// проверка на недопустимые символы
	var matched bool
	matched, _ = regexp.MatchString(`[^0-9,^\+,^\-,^\*,^\/,^I,^V,^X]`, a)

	if matched {
		fmt.Println("Строка не является математической операцией. Работа приложения завершена.")
		os.Exit(1)
	}

	// проверка на двойной операнд. кривенькое условие - от одного да 4х раз вызывается функция, но лучше не придумал
	if strings.Count(a, "+") > 1 || strings.Count(a, "-") > 1 || strings.Count(a, "*") > 1 || strings.Count(a, "/") > 1 {
		fmt.Println("Формат математической операции не удовлетворяет заданию. Работа приложения завершена.")
		os.Exit(1)
	}

	// проверка на отсутсвие операндов;)
	matched, _ = regexp.MatchString(`[\+,\-,\*,\/]`, a)
	
	if !matched {
		fmt.Println("В выражении отсутствуют операнды. Работа приложения завершена.")
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
		fmt.Println("Выражение должно состоять только из арабских или только из римских чисел, разделенных операндом. Работа приложения завершена.")
		os.Exit(1)
	}

	// разделяем строку по операнду, получаем операнд в виде символа
	rCipher := regexp.MustCompile(`[\+,\-,\*,\/]`)
	rOperand := regexp.MustCompile(`[0-9,I,V,X]`) 

	var op string
	for _, o := range rOperand.Split(a, -1) {
		if o != "" {
			op = o
			break
		}
	}

	var err error
	var d[2]int

	// цикл по двум числам, с записью в массив, и проверка на вес
	for n, b := range rCipher.Split(a, -1) {
			
		if b == "" {
			continue
		}

		if roman {
			d[n], err = romannumeral.StringToInt(b)
		}else{
			d[n], err = strconv.Atoi(b)	
		}
				
		if err != nil || d[n] > 10 || d[n] < 1 {
			fmt.Println("Формат математической операции не удовлетворяет заданию. Работа приложения завершена.")
			os.Exit(1)
		}
		
	}
		
	// вычисляем результат
	var res int
	switch op {
		case "+":
			res = d[0] + d[1]
		case "-":
			res = d[0] - d[1]
		case "*":
			res = d[0] * d[1]
		case "/":
			res = d[0] / d[1]
		}

	// проверка на отрицательные римские числа
	if roman && res < 1 {
		fmt.Println("Результат вычисления выражения в римских числах не может быть отрицательным. Работа приложения завершена.")
		os.Exit(1)
	}

	// Вывод результата
	if roman {
		var resR string
		resR, err = romannumeral.IntToString(res)
		if err != nil {
			fmt.Println("Неизвестная ошибка. Работа приложения завершена.")
			os.Exit(1)
		}
		fmt.Printf("Результат выполнения операции: %s\n", resR)
	}else{
		fmt.Printf("Результат выполнения операции: %d\n", res)
	}

	fmt.Println("Работа выполнена")

	os.Exit(1)

}
